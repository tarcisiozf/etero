package evm

type Opcode byte

var Instructions = make([]*Instruction, 0)
var InstructionsByOpcode = make(map[Opcode]*Instruction)

type Instruction struct {
	Opcode  Opcode
	Name    string
	Execute func(ctx *ExecutionContext)
}

func registerInstruction(opcode Opcode, name string, executeFunc func(ctx *ExecutionContext)) *Instruction {
	ix := &Instruction{
		Opcode:  opcode,
		Name:    name,
		Execute: executeFunc,
	}
	Instructions = append(Instructions, ix)

	if InstructionsByOpcode[opcode] != nil {
		panic("duplicated opcode")
	}
	InstructionsByOpcode[opcode] = ix

	return ix
}
