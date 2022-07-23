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
}

func newExecutionContext(code []byte) *ExecutionContext {
	return &ExecutionContext{
		code:       code,
		stack:      newStack(),
		memory:     newMemory(),
		returnData: make([]byte, 0),
	}
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
