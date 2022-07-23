package evm

import (
	"etero/evm/word"
	"fmt"
)

type Memory struct {
	storage map[string]byte
}

func newMemory() *Memory {
	return &Memory{
		storage: make(map[string]byte),
	}
}

func (mem *Memory) store(offset word.Word, value byte) {
	mem.storage[offset.String()] = value
}

func (mem *Memory) load(offset word.Word) byte {
	return mem.storage[offset.String()]
}

func (mem *Memory) loadRange(offset word.Word, length uint64) []byte {
	data := make([]byte, length)

	for i := uint64(0); i < length; i++ {
		pos := word.NewWord().Add(offset, word.NewFromInt(i))
		data[i] = mem.storage[pos.String()]
	}

	return data
}

func (mem *Memory) print() {
	items := make([]byte, len(mem.storage))
	i := 0

	for _, item := range mem.storage {
		items[i] = item
		i++
	}

	fmt.Println("memory:", items)
}

func (mem *Memory) loadWord(offset word.Word) word.Word {
	return word.NewFromBytes(mem.loadRange(offset, 32))
}

func (mem *Memory) storeWord(offset, value word.Word) {
	hex := []byte(value.String())[2:]
	for i := uint64(0); i < 32; i++ {
		pos := word.NewWord().Add(offset, word.NewFromInt(i))
		mem.store(pos, hex[i])
	}
}

func (mem *Memory) size() int {
	return len(mem.storage) // size of 32
}
