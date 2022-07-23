package evm

import (
	"etero/evm/word"
	"fmt"
)

type Memory struct {
	storage map[string]byte
}

func NewMemory() *Memory {
	return &Memory{
		storage: make(map[string]byte),
	}
}

func (mem *Memory) Store(offset word.Word, value byte) {
	mem.storage[offset.String()] = value
}

func (mem *Memory) Load(offset word.Word) byte {
	return mem.storage[offset.String()]
}

func (mem *Memory) LoadRange(offset word.Word, length uint64) []byte {
	data := make([]byte, length)

	for i := uint64(0); i < length; i++ {
		pos := word.NewWord().Add(offset, word.NewInt(i))
		data[i] = mem.storage[pos.String()]
	}

	return data
}

func (mem *Memory) Print() {
	items := make([]byte, len(mem.storage))
	i := 0

	for _, item := range mem.storage {
		items[i] = item
		i++
	}

	fmt.Println("memory:", items)
}
