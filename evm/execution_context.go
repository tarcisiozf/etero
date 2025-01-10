package evm

const StackMaxDepth = 1024 // TODO: remove from here

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
		code:   code,
		stack:  NewStack(StackMaxDepth),
		memory: NewMemory(),
	}
}

func (ec *ExecutionContext) stop() {
	ec.stopped = true
}

func (ec *ExecutionContext) readCode(numBytes int) []byte {
	slice := ec.code[ec.pc : ec.pc+numBytes]
	ec.pc += numBytes
	return slice
}

func (ec *ExecutionContext) setReturnData(offset, length Word) error {
	ec.stop()
	returnData, err := ec.memory.loadRange(offset, length)
	if err != nil {
		return err
	}
	ec.returnData = returnData
	return nil
}
