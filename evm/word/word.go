package word

import "github.com/holiman/uint256"

type Word = *uint256.Int

func NewWord() Word {
	return uint256.NewInt(0)
}

func NewInt(val uint64) Word {
	return uint256.NewInt(val)
}
