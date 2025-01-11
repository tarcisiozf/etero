package evm

const StackMaxDepth = 1024 // TODO: remove from here

type ExecutionContext struct {
	code       []byte
	stack      *Stack
	memory     *Memory
	pc         int // program counter
	stopped    bool
	returnData []byte
	jumpdests  map[int]bool
}

func NewExecutionContext(code []byte) *ExecutionContext {
	return &ExecutionContext{
		code:      code,
		stack:     NewStack(StackMaxDepth),
		memory:    NewMemory(),
		jumpdests: mapValidJumpDestinations(code),
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

func mapValidJumpDestinations(code []byte) map[int]bool {
	destinations := make(map[int]bool)
	i := 0
	for i < len(code) {
		currentOp := code[i]
		if currentOp == JumpDest.opcode {
			destinations[i] = true
		} else if Push1.opcode <= currentOp && currentOp <= Push32.opcode {
			i += int(currentOp-Push1.opcode) + 1
		}
		i++
	}
	return destinations
}
