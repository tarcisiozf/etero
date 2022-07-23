package evm

import "etero/evm/word"

var STOP = registerInstruction(0x00, "STOP", func(ctx *ExecutionContext) {
	ctx.stop()
})

var PUSH1 = registerInstruction(0x60, "PUSH1", func(ctx *ExecutionContext) {
	data := ctx.readCode(1)[0]
	conv := word.NewFromInt(uint64(data))
	ctx.stack.push(conv)
})

// TODO: implement until PUSH32

var PUSH32 = registerInstruction(0x7f, "PUSH32", func(ctx *ExecutionContext) {
	panic("not implemented")
})

var ADD = registerInstruction(0x01, "ADD", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	ctx.stack.push(word.NewWord().Add(a, b))
})

var MUL = registerInstruction(0x02, "MUL", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	ctx.stack.push(word.NewWord().Mul(a, b))
})

var SUB = registerInstruction(0x03, "SUB", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	ctx.stack.push(word.NewWord().Sub(a, b))
})

var LT = registerInstruction(0x10, "LT", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	var result uint64
	if a.Lt(b) {
		result = 1
	}
	ctx.stack.push(word.NewFromInt(result))
})

var GT = registerInstruction(0x11, "GT", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	var result uint64
	if a.Gt(b) {
		result = 1
	}
	ctx.stack.push(word.NewFromInt(result))
})

var EQ = registerInstruction(0x14, "EQ", func(ctx *ExecutionContext) {
	a := ctx.stack.pop()
	b := ctx.stack.pop()
	var result uint64
	if a.Eq(b) {
		result = 1
	}
	ctx.stack.push(word.NewFromInt(result))
})

var ISZERO = registerInstruction(0x15, "ISZERO", func(ctx *ExecutionContext) {
	item := ctx.stack.pop()
	var result uint64
	if item.IsZero() {
		result = 1
	}
	ctx.stack.push(word.NewFromInt(result))
})

var POP = registerInstruction(0x50, "POP", func(ctx *ExecutionContext) {
	ctx.stack.pop()
})

var MLOAD = registerInstruction(0x51, "MLOAD", func(ctx *ExecutionContext) {
	offset := ctx.stack.pop()
	w := ctx.memory.loadWord(offset)
	ctx.stack.push(w)
})

var MSTORE = registerInstruction(0x52, "MSTORE", func(ctx *ExecutionContext) {
	offset := ctx.stack.pop()
	value := ctx.stack.pop()
	ctx.memory.storeWord(offset, value)
})

var MSTORE8 = registerInstruction(0x53, "MSTORE8", func(ctx *ExecutionContext) {
	offset := ctx.stack.pop()
	value := ctx.stack.pop() // TODO: modulo 256
	ctx.memory.store(offset, byte(value.Uint64()))
})

var JUMP = registerInstruction(0x56, "JUMP", func(ctx *ExecutionContext) {
	targetDest := int(ctx.stack.pop().Uint64())
	ctx.jump(targetDest)
})

var JUMPI = registerInstruction(0x57, "JUMPI", func(ctx *ExecutionContext) {
	targetDest := int(ctx.stack.pop().Uint64())
	cond := ctx.stack.pop()

	if !cond.IsZero() {
		ctx.jump(targetDest)
	}
})

var PC = registerInstruction(0x58, "PC", func(ctx *ExecutionContext) {
	pc := word.NewFromInt(uint64(ctx.pc))
	ctx.stack.push(pc)
})

var MSIZE = registerInstruction(0x59, "MSIZE", func(ctx *ExecutionContext) {
	size := word.NewFromInt(32 * uint64(ctx.memory.size()))
	ctx.stack.push(size)
})

var JUMPDEST = registerInstruction(0x5b, "JUMPDEST", func(ctx *ExecutionContext) {})

var DUP1 = registerInstruction(0x80, "DUP1", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(1))
})

var DUP2 = registerInstruction(0x81, "DUP2", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(2))
})

var DUP3 = registerInstruction(0x82, "DUP3", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(3))
})

var DUP4 = registerInstruction(0x83, "DUP4", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(4))
})

var DUP5 = registerInstruction(0x84, "DUP5", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(5))
})

var DUP6 = registerInstruction(0x85, "DUP6", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(6))
})

var DUP7 = registerInstruction(0x86, "DUP7", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(7))
})

var DUP8 = registerInstruction(0x87, "DUP8", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(8))
})

var DUP9 = registerInstruction(0x88, "DUP9", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(9))
})

var DUP10 = registerInstruction(0x89, "DUP10", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(10))
})

var DUP11 = registerInstruction(0x8a, "DUP11", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(11))
})

var DUP12 = registerInstruction(0x8b, "DUP12", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(12))
})

var DUP13 = registerInstruction(0x8c, "DUP13", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(13))
})

var DUP14 = registerInstruction(0x8d, "DUP14", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(14))
})

var DUP15 = registerInstruction(0x8e, "DUP15", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(15))
})

var DUP16 = registerInstruction(0x8f, "DUP16", func(ctx *ExecutionContext) {
	ctx.stack.push(ctx.stack.peek(16))
})

var SWAP1 = registerInstruction(0x90, "SWAP1", func(ctx *ExecutionContext) {
	ctx.stack.swap(1)
})

var SWAP2 = registerInstruction(0x91, "SWAP2", func(ctx *ExecutionContext) {
	ctx.stack.swap(2)
})

var SWAP3 = registerInstruction(0x92, "SWAP3", func(ctx *ExecutionContext) {
	ctx.stack.swap(3)
})

var SWAP4 = registerInstruction(0x93, "SWAP4", func(ctx *ExecutionContext) {
	ctx.stack.swap(4)
})

var SWAP5 = registerInstruction(0x94, "SWAP5", func(ctx *ExecutionContext) {
	ctx.stack.swap(5)
})

var SWAP6 = registerInstruction(0x95, "SWAP6", func(ctx *ExecutionContext) {
	ctx.stack.swap(6)
})

var SWAP7 = registerInstruction(0x96, "SWAP7", func(ctx *ExecutionContext) {
	ctx.stack.swap(7)
})

var SWAP8 = registerInstruction(0x97, "SWAP8", func(ctx *ExecutionContext) {
	ctx.stack.swap(8)
})

var SWAP9 = registerInstruction(0x98, "SWAP9", func(ctx *ExecutionContext) {
	ctx.stack.swap(9)
})

var SWAP10 = registerInstruction(0x99, "SWAP10", func(ctx *ExecutionContext) {
	ctx.stack.swap(10)
})

var SWAP11 = registerInstruction(0x9A, "SWAP11", func(ctx *ExecutionContext) {
	ctx.stack.swap(11)
})

var SWAP12 = registerInstruction(0x9B, "SWAP12", func(ctx *ExecutionContext) {
	ctx.stack.swap(12)
})

var SWAP13 = registerInstruction(0x9C, "SWAP13", func(ctx *ExecutionContext) {
	ctx.stack.swap(13)
})

var SWAP14 = registerInstruction(0x9D, "SWAP14", func(ctx *ExecutionContext) {
	ctx.stack.swap(14)
})

var SWAP15 = registerInstruction(0x9E, "SWAP15", func(ctx *ExecutionContext) {
	ctx.stack.swap(15)
})

var SWAP16 = registerInstruction(0x9F, "SWAP16", func(ctx *ExecutionContext) {
	ctx.stack.swap(16)
})

var RETURN = registerInstruction(0xf3, "RETURN", func(ctx *ExecutionContext) {
	offset := ctx.stack.pop()
	length := ctx.stack.pop().Uint64()
	ctx.setReturnData(offset, length)
})

var REVERT = registerInstruction(0xfd, "REVERT", func(ctx *ExecutionContext) {
	// TODO: no-op for now
	ctx.stop()
})
