package evm

type Instruction struct {
	opcode   int
	name     string
	execFunc func(ctx *ExecutionContext) error
}
