package evm

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOpcode_Stop(t *testing.T) {
	ctx := NewExecutionContext(nil)
	err := Stop.execFunc(ctx)
	assert.Nil(t, err)
	assert.True(t, ctx.stopped)
}

func TestOpcode_Pushes(t *testing.T) {
	tests := []struct {
		ix   *Instruction
		size int
	}{
		{Push1, 1},
		{Push2, 2},
		{Push3, 3},
		{Push4, 4},
		{Push5, 5},
		{Push6, 6},
		{Push7, 7},
		{Push8, 8},
		{Push9, 9},
		{Push10, 10},
		{Push11, 11},
		{Push12, 12},
		{Push13, 13},
		{Push14, 14},
		{Push15, 15},
		{Push16, 16},
		{Push17, 17},
		{Push18, 18},
		{Push19, 19},
		{Push20, 20},
		{Push21, 21},
		{Push22, 22},
		{Push23, 23},
		{Push24, 24},
		{Push25, 25},
		{Push26, 26},
		{Push27, 27},
		{Push28, 28},
		{Push29, 29},
		{Push30, 30},
		{Push31, 31},
		{Push32, 32},
	}

	for _, test := range tests {
		t.Run(test.ix.name, func(t *testing.T) {
			code := make([]byte, test.size)
			_, _ = rand.Read(code)

			ctx := NewExecutionContext(code)
			err := test.ix.execFunc(ctx)
			assert.Nil(t, err)

			w, err := ctx.stack.pop()
			assert.Nil(t, err)
			assert.Equal(
				t,
				"0x"+strings.TrimLeft(hex.EncodeToString(code), "0"),
				w.internal.Hex(),
			)
		})
	}
}

func TestOpcode_Add(t *testing.T) {
	ctx := NewExecutionContext(nil)
	_ = ctx.stack.push(NewWordFromUint64(1))
	_ = ctx.stack.push(NewWordFromUint64(2))

	err := Add.execFunc(ctx)
	assert.Nil(t, err)
	w, err := ctx.stack.pop()
	assert.Nil(t, err)
	assert.Equal(t, uint64(3), w.Uint64())
}

func TestOpcode_Mul(t *testing.T) {
	ctx := NewExecutionContext(nil)
	_ = ctx.stack.push(NewWordFromUint64(2))
	_ = ctx.stack.push(NewWordFromUint64(3))

	err := Mul.execFunc(ctx)
	assert.Nil(t, err)
	w, err := ctx.stack.pop()
	assert.Nil(t, err)
	assert.Equal(t, uint64(6), w.Uint64())
}

func TestOpcode_Mstore8(t *testing.T) {
	ctx := NewExecutionContext(nil)
	offset := NewWordFromUint64(0)
	value := byte(42)
	_ = ctx.stack.push(NewWordFromUint64(uint64(value)))
	_ = ctx.stack.push(offset)

	err := Mstore8.execFunc(ctx)
	assert.Nil(t, err)
	b := ctx.memory.load(offset)
	assert.Equal(t, value, b)
}

func TestOpcode_Return(t *testing.T) {
	ctx := NewExecutionContext(nil)
	offset := NewWordFromUint64(0)
	length := 1
	value := byte(42)
	_ = ctx.stack.push(NewWordFromUint64(uint64(length)))
	_ = ctx.stack.push(offset)
	ctx.memory.store(offset, value)

	err := Return.execFunc(ctx)
	assert.Nil(t, err)
	assert.Equal(t, []byte{value}, ctx.returnData)
}
