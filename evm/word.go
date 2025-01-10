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

func (w Word) Mod(v Word) Word {
	return Word{uint256.NewInt(0).Mod(w.internal, v.internal)}
}

func (w Word) Uint64() uint64 {
	return w.internal.Uint64()
}

func NewWordFromUint64(value uint64) Word {
	return Word{uint256.NewInt(value)}
}
