package evm

import "etero/evm/word"

const MaxStackSize = 1024

type Stack struct {
	storage []word.Word
}

func NewStack() *Stack {
	return &Stack{
		storage: make([]word.Word, 0),
	}
}

func (st *Stack) Push(item word.Word) {
	if st.Size() == MaxStackSize {
		panic("stack overflow")
	}

	st.storage = append(st.storage, item)
}

func (st *Stack) Pop() word.Word {
	if st.Size() == 0 {
		panic("stack underflow")
	}

	item := st.storage[0]
	st.storage = st.storage[1:]

	return item
}

func (st *Stack) Size() int {
	return len(st.storage)
}
