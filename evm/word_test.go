package evm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWordFromUint64(t *testing.T) {
	value := uint64(42)
	w := NewWordFromUint64(value)
	assert.Equal(t, value, w.internal.Uint64())
}

func TestWord_String(t *testing.T) {
	value := uint64(42)
	w := NewWordFromUint64(value)
	assert.Equal(t, "42", w.String())
}

func TestWord_Uint64(t *testing.T) {
	value := uint64(42)
	w := NewWordFromUint64(value)
	assert.Equal(t, value, w.Uint64())
}

func TestWord_Add(t *testing.T) {
	a := NewWordFromUint64(1)
	b := NewWordFromUint64(2)
	c := a.Add(b)
	assert.Equal(t, uint64(3), c.Uint64())
}

func TestWord_Mul(t *testing.T) {
	a := NewWordFromUint64(2)
	b := NewWordFromUint64(3)
	c := a.Mul(b)
	assert.Equal(t, uint64(6), c.Uint64())
}

func TestWord_Mod(t *testing.T) {
	a := NewWordFromUint64(2)
	b := NewWordFromUint64(3)
	c := a.Mod(b)
	assert.Equal(t, uint64(2), c.Uint64())
}
