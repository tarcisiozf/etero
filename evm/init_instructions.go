package evm

import "etero/evm/word"

var STOP = registerInstruction(0x00, "STOP", func(execCtx *ExecutionContext) {
	execCtx.stop()
})

var PUSH1 = registerInstruction(0x60, "PUSH1", func(execCtx *ExecutionContext) {
	data := execCtx.readCode(1)[0]
	conv := word.NewInt(uint64(data))
	execCtx.stack.push(conv)
})

// TODO: implement until PUSH32

var PUSH32 = registerInstruction(0x7f, "PUSH32", func(execCtx *ExecutionContext) {
	panic("not implemented")
})

var ADD = registerInstruction(0x01, "ADD", func(execCtx *ExecutionContext) {
	a := execCtx.stack.pop()
	b := execCtx.stack.pop()
	execCtx.stack.push(word.NewWord().Add(a, b))
})

var MUL = registerInstruction(0x02, "MUL", func(execCtx *ExecutionContext) {
	a := execCtx.stack.pop()
	b := execCtx.stack.pop()
	execCtx.stack.push(word.NewWord().Mul(a, b))
})

var MSTORE8 = registerInstruction(0x53, "MSTORE8", func(execCtx *ExecutionContext) {
	offset := execCtx.stack.pop()
	data := execCtx.stack.pop() // TODO: modulo 256
	execCtx.memory.store(offset, byte(data.Uint64()))
})

var RETURN = registerInstruction(0xf3, "RETURN", func(execCtx *ExecutionContext) {
	offset := execCtx.stack.pop()
	length := execCtx.stack.pop().Uint64()
	execCtx.setReturnData(offset, length)
})

var JUMPDEST = registerInstruction(0x5b, "JUMPDEST", func(execCtx *ExecutionContext) {})

var JUMP = registerInstruction(0x56, "JUMP", func(execCtx *ExecutionContext) {
	targetDest := int(execCtx.stack.pop().Uint64())
	execCtx.jump(targetDest)
})

var JUMPI = registerInstruction(0x57, "JUMPI", func(execCtx *ExecutionContext) {
	targetDest := int(execCtx.stack.pop().Uint64())
	cond := execCtx.stack.pop()

	if !cond.IsZero() {
		execCtx.jump(targetDest)
	}
})

var PC = registerInstruction(0x58, "PC", func(execCtx *ExecutionContext) {
	pc := word.NewInt(uint64(execCtx.pc))
	execCtx.stack.push(pc)
})

var DUP1 = registerInstruction(0x80, "DUP1", func(execCtx *ExecutionContext) {
	execCtx.stack.push(execCtx.stack.peek(0))
})

var DUP2 = registerInstruction(0x81, "DUP2", func(execCtx *ExecutionContext) {
	execCtx.stack.push(execCtx.stack.peek(1))
})

var DUP3 = registerInstruction(0x82, "DUP3", func(execCtx *ExecutionContext) {
	execCtx.stack.push(execCtx.stack.peek(2))
})

// TODO: Implement until DUP16
