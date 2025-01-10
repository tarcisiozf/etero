package evm

type Memory struct {
	storage map[int]byte
}

func NewMemory() *Memory {
	return &Memory{
		storage: make(map[int]byte, 0),
	}
}

func (m *Memory) store(offset int, value byte) {
	m.storage[offset] = value
}

func (m *Memory) load(offset int) byte {
	return m.storage[offset]
}
