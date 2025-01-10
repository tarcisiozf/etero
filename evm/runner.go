package evm

import (
	"errors"
	"fmt"
)

func Run(code []byte) error {
	execContext := NewExecutionContext(code)
	instructionRegistry := NewInstructionsRegistry()

	err := registerEvmOpcodes(instructionRegistry)
	if err != nil {
		return err
	}

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
	if execContext.pc < 0 {
		return nil, errors.New("invalid code offset")
	}

	// section 9.4.1 of the yellow paper, the operation to be executed if pc is outside code is STOP
	if execContext.pc >= len(execContext.code) {
		return Stop, nil
	}

	opcode := int(execContext.readCode(1)[0])
	instruction := registry.find(opcode)
	if instruction == nil {
		return nil, fmt.Errorf("unknow opcode %x", opcode)
	}

	return instruction, nil
}
