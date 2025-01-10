package evm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExecutionContext(t *testing.T) {
	code := []byte{1, 2, 3, 4, 5, 6}
	execCtx := NewExecutionContext(code)

	t.Run("stop", func(t *testing.T) {
		assert.False(t, execCtx.stopped)
		execCtx.stop()
		assert.True(t, execCtx.stopped)
	})

	// TODO: refactor using test table
	t.Run("readCode", func(t *testing.T) {
		slice := execCtx.readCode(3)
		assert.Equal(t, []byte{1, 2, 3}, slice)

		// check second pass to see if cursor moved
		slice = execCtx.readCode(1)
		assert.Equal(t, []byte{4}, slice)
	})
}
