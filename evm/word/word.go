package word

import "github.com/holiman/uint256"

type Word = *uint256.Int

func NewWord() Word {
	return uint256.NewInt(0)
}

func NewFromInt(val uint64) Word {
	return uint256.NewInt(val)
}

func NewFromBytes(b []byte) Word {
	hex := "0x" + string(b)
	w, err := uint256.FromHex(hex)
	if err != nil {
		panic(err)
	}
	return w
}
