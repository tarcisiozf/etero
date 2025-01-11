package evm

type Instruction struct {
	opcode   byte
	name     string
	execFunc func(ctx *ExecutionContext) error
}
