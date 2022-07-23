package evm

import (
	"fmt"
	"strings"
)

type ExecutionContext struct {
	code    []byte
	stack   *Stack
	memory  *Memory
	pc      int
	stopped bool
}

func NewExecutionContext(code []byte) *ExecutionContext {
	return &ExecutionContext{
		code:   code,
		stack:  NewStack(),
		memory: NewMemory(),
	}
}

func (execCtx *ExecutionContext) Stop() {
	execCtx.stopped = true
}

func (execCtx *ExecutionContext) IsRunning() bool {
	return !execCtx.stopped
}

func (execCtx *ExecutionContext) Pc() int {
	return execCtx.pc
}

func (execCtx *ExecutionContext) ReadCode(numBytes int) []byte {
	slice := execCtx.code[execCtx.pc : execCtx.pc+numBytes]
	execCtx.pc += numBytes
	return slice
}

func (execCtx *ExecutionContext) PrintStack() {
	items := make([]string, len(execCtx.stack.storage))

	for i, item := range execCtx.stack.storage {
		items[i] = item.String()
	}

	fmt.Printf("STACK: [%s]\n", strings.Join(items, " "))
}
