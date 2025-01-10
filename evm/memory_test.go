package evm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMemory(t *testing.T) {
	memory := NewMemory()
	offset := 0
	value := byte(42)

	t.Run("store", func(t *testing.T) {
		assert.Equal(t, 0, len(memory.storage))
		memory.store(offset, value)
		assert.Equal(t, 1, len(memory.storage))
	})

	t.Run("load", func(t *testing.T) {
		v := memory.load(offset)
		assert.Equal(t, value, v)

		t.Run("loading empty offset defaults to zero", func(t *testing.T) {
			v = memory.load(-1)
			assert.Equal(t, byte(0), v)
		})
	})
}
