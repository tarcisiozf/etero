package evm

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: improve test coverage
func TestRun(t *testing.T) {
	code, err := hex.DecodeString("600660070200")
	assert.Nil(t, err)

	err = Run(code)
	assert.Nil(t, err)
}
