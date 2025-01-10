package evm

import (
	"errors"
	"fmt"
)

func Run(code []byte) error {
	execContext := NewExecutionContext(code)
	instructionRegistry := NewInstructionsRegistry()

	registerEvmOpcodes(instructionRegistry)

	for !execContext.stopped {
		instruction, err := decodeOpcode(execContext, instructionRegistry)
		if err != nil {
			return err
		}
		err = instruction.execFunc(execContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func registerEvmOpcodes(registry *InstructionsRegistry) error {
	instructions := []*Instruction{Stop, Add, Mul, Push1}
	for _, ix := range instructions {
		if err := registry.register(ix); err != nil {
			return err
		}
	}
	return nil
}

func decodeOpcode(execContext *ExecutionContext, registry *InstructionsRegistry) (*Instruction, error) {
	if execContext.pc < 0 || execContext.pc > len(execContext.code) {
		return nil, errors.New("invalid code offset")
	}

	opcode := int(execContext.readCode(1)[0])
	instruction := registry.find(opcode)
	if instruction == nil {
		return nil, fmt.Errorf("unknow opcode %x", opcode)
	}

	return instruction, nil
}
