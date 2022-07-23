package evm

import (
	"etero/evm/exec"
	"fmt"
)

func Run(code []byte) {
	ctx := exec.NewExecutionContext(code)

	for ctx.IsRunning() {
		prevPc := ctx.Pc()
		ix := decodeOpcode(ctx)
		ix.Execute(ctx)

		fmt.Printf("ix: %s | pc: %d\n", ix.Name, prevPc)
		ctx.PrintStack()
	}
}

func decodeOpcode(execCtx *exec.ExecutionContext) *exec.Instruction {
	opcode := exec.Opcode(execCtx.ReadCode(1)[0])
	ix := exec.InstructionsByOpcode[opcode]
	if ix == nil {
		panic("unknown opcode")
	}

	return ix
}
