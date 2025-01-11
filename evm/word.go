package evm

import (
	"encoding/hex"
	"fmt"
	"github.com/holiman/uint256"
	"strings"
)

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

func NewWordFromBytes(b []byte) Word {
	size := len(b)
	if size == 0 || size > 32 {
		panic(fmt.Sprintf("invalid word size: %d", size))
	}
	h := "0x" + strings.TrimLeft(hex.EncodeToString(b), "0")
	return Word{uint256.MustFromHex(h)}
}
