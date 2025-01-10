package evm

var (
	Stop = &Instruction{
		opcode: 0x00,
		name:   "STOP",
		execFunc: func(ctx *ExecutionContext) error {
			ctx.stop()
			return nil
		},
	}

	Push1 = &Instruction{
		opcode: 0x60,
		name:   "PUSH1",
		execFunc: func(ctx *ExecutionContext) error {
			data := ctx.readCode(1)[0]
			word := NewWordFromUint64(uint64(data))
			return ctx.stack.push(word)
		},
	}

	Add = &Instruction{
		opcode: 0x01,
		name:   "ADD",
		execFunc: func(ctx *ExecutionContext) error {
			a, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			b, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			return ctx.stack.push(a.Add(b)) // TODO: mod 2**256
		},
	}

	Mul = &Instruction{
		opcode: 0x02,
		name:   "MUL",
		execFunc: func(ctx *ExecutionContext) error {
			a, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			b, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			return ctx.stack.push(a.Mul(b)) // TODO: mod 2**256
		},
	}
)
