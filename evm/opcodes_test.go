package evm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpcode_Stop(t *testing.T) {
	ctx := NewExecutionContext(nil)
	err := Stop.execFunc(ctx)
	assert.Nil(t, err)
	assert.True(t, ctx.stopped)
}

func TestOpcode_Push1(t *testing.T) {
	ctx := NewExecutionContext([]byte{1})
	err := Push1.execFunc(ctx)
	assert.Nil(t, err)
	w, err := ctx.stack.pop()
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), w.Uint64())
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
