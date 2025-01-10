package evm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMemory(t *testing.T) {
	memory := NewMemory()
	offset := NewWordFromUint64(0)
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
			offset := NewWordFromUint64(33)
			v = memory.load(offset)
			assert.Equal(t, byte(0), v)
		})
	})

	t.Run("loadRange", func(t *testing.T) {
		length := uint64(3)
		for i := uint64(0); i < length; i++ {
			memory.store(NewWordFromUint64(i), byte(i)+1)
		}

		data, err := memory.loadRange(offset, NewWordFromUint64(length))
		assert.Nil(t, err)
		assert.Equal(t, []byte{1, 2, 3}, data)
	})
}
