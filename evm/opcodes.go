package evm

import "errors"

var (
	Stop = &Instruction{
		opcode: 0x00,
		name:   "STOP",
		execFunc: func(ctx *ExecutionContext) error {
			ctx.stop()
			return nil
		},
	}

	Push1  = makePushFunc(0x60, "PUSH1", 1)
	Push2  = makePushFunc(0x61, "PUSH2", 2)
	Push3  = makePushFunc(0x62, "PUSH3", 3)
	Push4  = makePushFunc(0x63, "PUSH4", 4)
	Push5  = makePushFunc(0x64, "PUSH5", 5)
	Push6  = makePushFunc(0x65, "PUSH6", 6)
	Push7  = makePushFunc(0x66, "PUSH7", 7)
	Push8  = makePushFunc(0x67, "PUSH8", 8)
	Push9  = makePushFunc(0x68, "PUSH9", 9)
	Push10 = makePushFunc(0x69, "PUSH10", 10)
	Push11 = makePushFunc(0x6A, "PUSH11", 11)
	Push12 = makePushFunc(0x6B, "PUSH12", 12)
	Push13 = makePushFunc(0x6C, "PUSH13", 13)
	Push14 = makePushFunc(0x6D, "PUSH14", 14)
	Push15 = makePushFunc(0x6E, "PUSH15", 15)
	Push16 = makePushFunc(0x6F, "PUSH16", 16)
	Push17 = makePushFunc(0x70, "PUSH17", 17)
	Push18 = makePushFunc(0x71, "PUSH18", 18)
	Push19 = makePushFunc(0x72, "PUSH19", 19)
	Push20 = makePushFunc(0x73, "PUSH20", 20)
	Push21 = makePushFunc(0x74, "PUSH21", 21)
	Push22 = makePushFunc(0x75, "PUSH22", 22)
	Push23 = makePushFunc(0x76, "PUSH23", 23)
	Push24 = makePushFunc(0x77, "PUSH24", 24)
	Push25 = makePushFunc(0x78, "PUSH25", 25)
	Push26 = makePushFunc(0x79, "PUSH26", 26)
	Push27 = makePushFunc(0x7A, "PUSH27", 27)
	Push28 = makePushFunc(0x7B, "PUSH28", 28)
	Push29 = makePushFunc(0x7C, "PUSH29", 29)
	Push30 = makePushFunc(0x7D, "PUSH30", 30)
	Push31 = makePushFunc(0x7E, "PUSH31", 31)
	Push32 = makePushFunc(0x7F, "PUSH32", 32)

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

	Sub = &Instruction{
		opcode: 0x03,
		name:   "SUB",
		execFunc: func(ctx *ExecutionContext) error {
			a, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			b, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			return ctx.stack.push(a.Sub(b)) // TODO: mod 2**256
		},
	}

	Mstore8 = &Instruction{
		opcode: 0x53,
		name:   "MSTORE8",
		execFunc: func(ctx *ExecutionContext) error {
			offset, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			value, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			ctx.memory.store(offset, byte(value.Mod(NewWordFromUint64(256)).Uint64()))
			return nil
		},
	}

	Return = &Instruction{
		opcode: 0xf3,
		name:   "RETURN",
		execFunc: func(ctx *ExecutionContext) error {
			offset, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			length, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			return ctx.setReturnData(offset, length)
		},
	}

	Jump = &Instruction{
		opcode: 0x56,
		name:   "JUMP",
		execFunc: func(ctx *ExecutionContext) error {
			dest, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			return doJump(ctx, int(dest.Uint64()))
		},
	}

	Jumpi = &Instruction{
		opcode: 0x57,
		name:   "JUMPI",
		execFunc: func(ctx *ExecutionContext) error {
			dest, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			cond, err := ctx.stack.pop()
			if err != nil {
				return err
			}
			if cond.Uint64() != 0 {
				return doJump(ctx, int(dest.Uint64()))
			}
			return nil
		},
	}

	JumpDest = &Instruction{
		opcode: 0x5b,
		name:   "JUMPDEST",
	}

	Pc = &Instruction{
		opcode: 0x58,
		name:   "PC",
		execFunc: func(ctx *ExecutionContext) error {
			pc := NewWordFromUint64(uint64(ctx.pc))
			return ctx.stack.push(pc)
		},
	}
)

func doJump(ctx *ExecutionContext, dest int) error {
	if !ctx.jumpdests[dest] {
		return errors.New("invalid jump destination")
	}
	ctx.pc = dest
	return nil
}

func makePushFunc(opcode byte, name string, numBytes int) *Instruction {
	return &Instruction{
		opcode: opcode,
		name:   name,
		execFunc: func(ctx *ExecutionContext) error {
			data := ctx.readCode(numBytes)
			word := NewWordFromBytes(data)
			return ctx.stack.push(word)
		},
	}
}
