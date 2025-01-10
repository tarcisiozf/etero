package evm

import "github.com/holiman/uint256"

type Word struct {
	internal *uint256.Int
}

func (w Word) String() string {
	return w.internal.String()
}

func NewWordFromUint64(value uint64) Word {
	return Word{uint256.NewInt(value)}
}
