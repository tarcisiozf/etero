package evm

import (
	"errors"
	"fmt"
)

var (
	Stop = &Instruction{
		opcode: 0x00,
		name:   "STOP",
		execFunc: func(ctx *ExecutionContext) error {
			ctx.stop()
			return nil
		},
	}

	Push1  = makePushFunc(0x60, 1)
	Push2  = makePushFunc(0x61, 2)
	Push3  = makePushFunc(0x62, 3)
	Push4  = makePushFunc(0x63, 4)
	Push5  = makePushFunc(0x64, 5)
	Push6  = makePushFunc(0x65, 6)
	Push7  = makePushFunc(0x66, 7)
	Push8  = makePushFunc(0x67, 8)
	Push9  = makePushFunc(0x68, 9)
	Push10 = makePushFunc(0x69, 10)
	Push11 = makePushFunc(0x6A, 11)
	Push12 = makePushFunc(0x6B, 12)
	Push13 = makePushFunc(0x6C, 13)
	Push14 = makePushFunc(0x6D, 14)
	Push15 = makePushFunc(0x6E, 15)
	Push16 = makePushFunc(0x6F, 16)
	Push17 = makePushFunc(0x70, 17)
	Push18 = makePushFunc(0x71, 18)
	Push19 = makePushFunc(0x72, 19)
	Push20 = makePushFunc(0x73, 20)
	Push21 = makePushFunc(0x74, 21)
	Push22 = makePushFunc(0x75, 22)
	Push23 = makePushFunc(0x76, 23)
	Push24 = makePushFunc(0x77, 24)
	Push25 = makePushFunc(0x78, 25)
	Push26 = makePushFunc(0x79, 26)
	Push27 = makePushFunc(0x7A, 27)
	Push28 = makePushFunc(0x7B, 28)
	Push29 = makePushFunc(0x7C, 29)
	Push30 = makePushFunc(0x7D, 30)
	Push31 = makePushFunc(0x7E, 31)
	Push32 = makePushFunc(0x7F, 32)

	Dup1  = makeDupFunc(0x80, 1)
	Dup2  = makeDupFunc(0x81, 2)
	Dup3  = makeDupFunc(0x82, 3)
	Dup4  = makeDupFunc(0x83, 4)
	Dup5  = makeDupFunc(0x84, 5)
	Dup6  = makeDupFunc(0x85, 6)
	Dup7  = makeDupFunc(0x86, 7)
	Dup8  = makeDupFunc(0x87, 8)
	Dup9  = makeDupFunc(0x88, 9)
	Dup10 = makeDupFunc(0x89, 10)
	Dup11 = makeDupFunc(0x8A, 11)
	Dup12 = makeDupFunc(0x8B, 12)
	Dup13 = makeDupFunc(0x8C, 13)
	Dup14 = makeDupFunc(0x8D, 14)
	Dup15 = makeDupFunc(0x8E, 15)
	Dup16 = makeDupFunc(0x8F, 16)

	Swap1  = makeSwapFunc(0x90, 1)
	Swap2  = makeSwapFunc(0x91, 2)
	Swap3  = makeSwapFunc(0x92, 3)
	Swap4  = makeSwapFunc(0x93, 4)
	Swap5  = makeSwapFunc(0x94, 5)
	Swap6  = makeSwapFunc(0x95, 6)
	Swap7  = makeSwapFunc(0x96, 7)
	Swap8  = makeSwapFunc(0x97, 8)
	Swap9  = makeSwapFunc(0x98, 9)
	Swap10 = makeSwapFunc(0x99, 10)
	Swap11 = makeSwapFunc(0x9A, 11)
	Swap12 = makeSwapFunc(0x9B, 12)
	Swap13 = makeSwapFunc(0x9C, 13)
	Swap14 = makeSwapFunc(0x9D, 14)
	Swap15 = makeSwapFunc(0x9E, 15)
	Swap16 = makeSwapFunc(0x9F, 16)

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

func makePushFunc(opcode byte, numBytes int) *Instruction {
	name := fmt.Sprintf("PUSH%d", numBytes)
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

func makeDupFunc(opcode byte, pos int) *Instruction {
	name := fmt.Sprintf("DUP%d", pos)
	return &Instruction{
		opcode: opcode,
		name:   name,
		execFunc: func(ctx *ExecutionContext) error {
			return ctx.stack.push(ctx.stack.peek(pos - 1))
		},
	}
}

func makeSwapFunc(opcode byte, pos int) *Instruction {
	name := fmt.Sprintf("SWAP%d", pos)
	return &Instruction{
		opcode: opcode,
		name:   name,
		execFunc: func(ctx *ExecutionContext) error {
			return ctx.stack.swap(pos - 1)
		},
	}
}
