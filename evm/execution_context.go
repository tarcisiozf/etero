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

func NewExecutionContext(code []byte) *ExecutionContext {
	return &ExecutionContext{
		code:       code,
		stack:      NewStack(),
		memory:     NewMemory(),
		returnData: make([]byte, 0),
	}
}

func (execCtx *ExecutionContext) Stop() {
	execCtx.stopped = true
}

func (execCtx *ExecutionContext) ReadCode(numBytes int) []byte {
	slice := execCtx.code[execCtx.pc : execCtx.pc+numBytes]
	execCtx.pc += numBytes
	return slice
}

func (execCtx *ExecutionContext) setReturnData(offset word.Word, length uint64) {
	execCtx.Stop()
	execCtx.returnData = execCtx.memory.LoadRange(offset, length)
}
