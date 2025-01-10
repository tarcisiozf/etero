package evm

import "github.com/holiman/uint256"

type Word struct {
	internal *uint256.Int
}

func (w Word) String() string {
	return w.internal.String()
}

func (w Word) Add(b Word) Word {
	return Word{uint256.NewInt(0).Add(w.internal, b.internal)}
}

func (w Word) Mul(b Word) Word {
	return Word{uint256.NewInt(0).Mul(w.internal, b.internal)}
}

func NewWordFromUint64(value uint64) Word {
	return Word{uint256.NewInt(value)}
}
