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

func (execCtx *ExecutionContext) stop() {
	execCtx.stopped = true
}

func (execCtx *ExecutionContext) readCode(numBytes int) []byte {
	slice := execCtx.code[execCtx.pc : execCtx.pc+numBytes]
	execCtx.pc += numBytes
	return slice
}

func (execCtx *ExecutionContext) setReturnData(offset word.Word, length uint64) {
	execCtx.stop()
	execCtx.returnData = execCtx.memory.loadRange(offset, length)
}

func (execCtx *ExecutionContext) jump(targetDest int) {
	if !execCtx.jumpdests[targetDest] {
		panic("invalid jump destination")
	}

	execCtx.pc = targetDest
}
