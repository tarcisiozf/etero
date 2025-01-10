package evm

import "errors"

type Stack struct {
	storage  []Word
	maxDepth int
}

func NewStack(maxDepth int) *Stack {
	return &Stack{
		storage:  make([]Word, 0),
		maxDepth: maxDepth,
	}
}

func (s *Stack) push(item Word) error {
	if len(s.storage)+1 > s.maxDepth {
		return errors.New("stack overflow")
	}

	s.storage = append(s.storage, item)
	return nil
}

func (s *Stack) pop() (w Word, e error) {
	size := len(s.storage)
	if size == 0 {
		return w, errors.New("stack underflow")
	}

	pos := size - 1
	item := s.storage[pos]
	s.storage = s.storage[:pos]

	return item, nil
}
