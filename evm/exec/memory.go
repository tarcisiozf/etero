package exec

import "etero/evm/word"

type Memory struct {
	storage map[word.Word]byte
}

func NewMemory() *Memory {
	return &Memory{
		storage: make(map[word.Word]byte),
	}
}

func (mem *Memory) Store(offset word.Word, value byte) {
	mem.storage[offset] = value
}

func (mem *Memory) Load(offset word.Word) byte {
	return mem.storage[offset]
}
