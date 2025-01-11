package evm

import (
	"fmt"
)

type InstructionsRegistry struct {
	instructions []*Instruction
	opcodeIndex  map[byte]*Instruction
}

func NewInstructionsRegistry() *InstructionsRegistry {
	return &InstructionsRegistry{
		instructions: make([]*Instruction, 0),
		opcodeIndex:  make(map[byte]*Instruction),
	}
}

func (ir *InstructionsRegistry) register(instruction *Instruction) error {
	if _, ok := ir.opcodeIndex[instruction.opcode]; ok {
		return fmt.Errorf("duplicated instruction with opcode %x", instruction.opcode)
	}

	ir.instructions = append(ir.instructions, instruction)
	ir.opcodeIndex[instruction.opcode] = instruction
	return nil
}

func (ir *InstructionsRegistry) find(opcode byte) *Instruction {
	return ir.opcodeIndex[opcode]
}
