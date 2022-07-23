package evm

import (
	"fmt"
)

func Run(code []byte) {
	ctx := newExecutionContext(code)

	for !ctx.stopped {
		prevPc := ctx.pc
		ix := decodeOpcode(ctx)
		ix.Execute(ctx)

		fmt.Printf("%s @ pc=%d\n", ix.Name, prevPc)
		ctx.stack.print()
		ctx.memory.print()
		fmt.Println()
	}

	fmt.Println("Output:", ctx.returnData)
}

func decodeOpcode(execCtx *ExecutionContext) *Instruction {
	// section 9.4.1 of the yellow paper, the operation to be executed if pc is outside code is STOP
	if execCtx.pc >= len(execCtx.code) {
		return STOP
	}

	opcode := Opcode(execCtx.readCode(1)[0])
	ix := InstructionsByOpcode[opcode]
	if ix == nil {
		panic(fmt.Sprintf("unknown opcode 0x%x", opcode))
	}

	return ix
}
