package evm

import (
	"etero/evm/word"
)

type ExecutionContext struct {
	code       []byte
	stack      *Stack
	memory     *Memory
	pc         int
	stopped    bool
	returnData []byte
	jumpdests  map[int]bool
}

func newExecutionContext(code []byte) *ExecutionContext {
	jumpdests := findValidJumpDestinations(code)

	return &ExecutionContext{
		code:       code,
		stack:      newStack(),
		memory:     newMemory(),
		returnData: make([]byte, 0),
		jumpdests:  jumpdests,
	}
}

func findValidJumpDestinations(code []byte) map[int]bool {
	dests := make(map[int]bool)
	i := 0

	for i < len(code) {
		currentOp := Opcode(code[i])
		if currentOp == JUMPDEST.Opcode {
			dests[i] = true
		} else if currentOp >= PUSH1.Opcode && currentOp <= PUSH32.Opcode {
			i += int(currentOp - PUSH1.Opcode + 1)
		}
		i++
	}

	return dests
}

func (ctx *ExecutionContext) stop() {
	ctx.stopped = true
}

func (ctx *ExecutionContext) readCode(numBytes int) []byte {
	slice := ctx.code[ctx.pc : ctx.pc+numBytes]
	ctx.pc += numBytes
	return slice
}

func (ctx *ExecutionContext) setReturnData(offset word.Word, length uint64) {
	ctx.stop()
	ctx.returnData = ctx.memory.loadRange(offset, length)
}

func (ctx *ExecutionContext) jump(targetDest int) {
	if !ctx.jumpdests[targetDest] {
		panic("invalid jump destination")
	}

	ctx.pc = targetDest
}
