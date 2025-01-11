package evm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	maxDepth := 1
	value := NewWordFromUint64(42)
	stack := NewStack(maxDepth)

	t.Run("push", func(t *testing.T) {
		err := stack.push(value)
		assert.Nil(t, err)

		t.Run("check overflow", func(t *testing.T) {
			err = stack.push(NewWordFromUint64(123))
			assert.NotNil(t, err)
			assert.Error(t, err, "stack overflow")
		})
	})

	t.Run("pop", func(t *testing.T) {
		item, err := stack.pop()
		assert.Nil(t, err)
		assert.Equal(t, value.String(), item.String())

		t.Run("check underflow", func(t *testing.T) {
			_, err = stack.pop()
			assert.NotNil(t, err)
			assert.Error(t, err, "stack underflow")
		})
	})

	t.Run("peek", func(t *testing.T) {
		_ = stack.push(value)
		item := stack.peek(0)
		assert.Equal(t, value.String(), item.String())
	})
}
