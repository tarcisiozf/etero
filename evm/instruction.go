package evm

type Instruction struct {
	opcode   int
	name     string
	execFund func(ctx *ExecutionContext) error
}
