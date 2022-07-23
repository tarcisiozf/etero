package evm

import (
	"fmt"
)

func Run(code []byte) {
	ctx := NewExecutionContext(code)

	for !ctx.stopped {
		prevPc := ctx.pc
		ix := decodeOpcode(ctx)
		ix.Execute(ctx)

		fmt.Printf("ix: %s | pc: %d\n", ix.Name, prevPc)
		ctx.PrintStack()
	}
}

func decodeOpcode(execCtx *ExecutionContext) *Instruction {
	opcode := Opcode(execCtx.ReadCode(1)[0])
	ix := InstructionsByOpcode[opcode]
	if ix == nil {
		panic("unknown opcode")
	}

	return ix
}
