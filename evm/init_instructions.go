package evm

import "etero/evm/word"

var STOP = registerInstruction(0x00, "STOP", func(execCtx *ExecutionContext) {
	execCtx.Stop()
})

var PUSH1 = registerInstruction(0x60, "PUSH1", func(execCtx *ExecutionContext) {
	data := execCtx.ReadCode(1)[0]
	conv := word.NewInt(uint64(data))
	execCtx.stack.Push(conv)
})

var ADD = registerInstruction(0x01, "ADD", func(execCtx *ExecutionContext) {
	a := execCtx.stack.Pop()
	b := execCtx.stack.Pop()
	execCtx.stack.Push(word.NewWord().Add(a, b))
})

var MUL = registerInstruction(0x02, "MUL", func(execCtx *ExecutionContext) {
	a := execCtx.stack.Pop()
	b := execCtx.stack.Pop()
	execCtx.stack.Push(word.NewWord().Mul(a, b))
})

var MSTORE8 = registerInstruction(0x53, "MSTORE8", func(execCtx *ExecutionContext) {
	offset := execCtx.stack.Pop()
	data := execCtx.stack.Pop() // TODO: modulo 256
	execCtx.memory.Store(offset, byte(data.Uint64()))
})

var RETURN = registerInstruction(0xf3, "RETURN", func(execCtx *ExecutionContext) {
	offset := execCtx.stack.Pop()
	length := execCtx.stack.Pop().Uint64()
	execCtx.setReturnData(offset, length)
})
