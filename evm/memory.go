package evm

import "errors"

type Memory struct {
	storage map[string]byte
}

func NewMemory() *Memory {
	return &Memory{
		storage: make(map[string]byte, 0),
	}
}

func (m *Memory) store(offset Word, value byte) {
	m.storage[offset.String()] = value
}

func (m *Memory) load(offset Word) byte {
	return m.storage[offset.String()]
}

func (m *Memory) loadRange(offset, length Word) ([]byte, error) {
	ioff := offset.Uint64()
	if ioff < 0 {
		return nil, errors.New("invalid memory access")
	}
	ilen := length.Uint64()
	data := make([]byte, ilen)
	for i := uint64(0); i < ilen; i++ {
		widx := NewWordFromUint64(i)
		data[i] = m.load(offset.Add(widx))
	}
	return data, nil
}
