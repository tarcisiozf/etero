package exec

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
