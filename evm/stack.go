package evm

import (
	"etero/evm/word"
	"fmt"
)

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
	if len(st.storage) == MaxStackSize {
		panic("stack overflow")
	}

	st.storage = append(st.storage, item)
}

func (st *Stack) Pop() word.Word {
	size := len(st.storage)
	if size == 0 {
		panic("stack underflow")
	}

	pos := size - 1
	item := st.storage[pos]
	st.storage = st.storage[:pos]

	return item
}

func (st *Stack) Print() {
	items := make([]string, len(st.storage))

	for i, item := range st.storage {
		items[i] = item.String()
	}

	fmt.Println("stack:", items)
}
