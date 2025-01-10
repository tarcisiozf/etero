package evm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInstructionsRegistry(t *testing.T) {
	ix := &Instruction{
		opcode: 0x42,
		name:   "test",
	}
	registry := NewInstructionsRegistry()

	t.Run("register", func(t *testing.T) {
		err := registry.register(ix)
		assert.Nil(t, err)

		t.Run("check for duplicates", func(t *testing.T) {
			err = registry.register(ix)
			assert.Error(t, err, "duplicated instruction with opcode 0x42")
		})
	})
}
