package evm

import "errors"

type Stack struct {
	storage  []int
	maxDepth int
}

func NewStack(maxDepth int) *Stack {
	return &Stack{
		storage:  make([]int, 0),
		maxDepth: maxDepth,
	}
}

func (s *Stack) push(item int) error {
	if len(s.storage)+1 > s.maxDepth {
		return errors.New("stack overflow")
	}

	s.storage = append(s.storage, item)
	return nil
}

func (s *Stack) pop() (int, error) {
	size := len(s.storage)
	if size == 0 {
		return 0, errors.New("stack underflow")
	}

	pos := size - 1
	item := s.storage[pos]
	s.storage = s.storage[:pos]

	return item, nil
}
