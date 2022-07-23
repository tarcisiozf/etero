package evm

import "etero/evm/word"

var STOP = registerInstruction(0x00, "STOP", func(ctx *ExecutionContext) {
	ctx.stop()
})

var PUSH1 = registerInstruction(0x60, "PUSH1", func(ctx *ExecutionContext) {
	data := ctx.readCode(1)[0]
	conv := word.NewInt(uint64(data))
	ctx.stack.push(conv)
})

// TODO: implement until PUSH32

var PUSH32 = registerInstruction(0x7f, "PUSH32", func(ctx *ExecutionContext) {
	panic("not implemented")
})

var ADD = registerInstruction(0x01, "ADD", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	ctx.stack.push(word.NewWord().Add(a, b))
})

var MUL = registerInstruction(0x02, "MUL", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	ctx.stack.push(word.NewWord().Mul(a, b))
})

var MSTORE8 = registerInstruction(0x53, "MSTORE8", func(ctx *ExecutionContext) {
	offset := ctx.stack.pop()
	data := ctx.stack.pop() // TODO: modulo 256
	ctx.memory.store(offset, byte(data.Uint64()))
})

var RETURN = registerInstruction(0xf3, "RETURN", func(ctx *ExecutionContext) {
	offset := ctx.stack.pop()
	length := ctx.stack.pop().Uint64()
	ctx.setReturnData(offset, length)
})

var JUMPDEST = registerInstruction(0x5b, "JUMPDEST", func(ctx *ExecutionContext) {})

var JUMP = registerInstruction(0x56, "JUMP", func(ctx *ExecutionContext) {
	targetDest := int(ctx.stack.pop().Uint64())
	ctx.jump(targetDest)
})

var JUMPI = registerInstruction(0x57, "JUMPI", func(ctx *ExecutionContext) {
	targetDest := int(ctx.stack.pop().Uint64())
	cond := ctx.stack.pop()

	if !cond.IsZero() {
		ctx.jump(targetDest)
	}
})

var PC = registerInstruction(0x58, "PC", func(ctx *ExecutionContext) {
	pc := word.NewInt(uint64(ctx.pc))
	ctx.stack.push(pc)
})

var DUP1 = registerInstruction(0x80, "DUP1", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(0))
})

var DUP2 = registerInstruction(0x81, "DUP2", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(1))
})

var DUP3 = registerInstruction(0x82, "DUP3", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(2))
})

// TODO: Implement until DUP16

var SWAP1 = registerInstruction(0x90, "SWAP1", func(ctx *ExecutionContext) {
	ctx.stack.swap(1)
})

// TODO: Implement until SWAP16
