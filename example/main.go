package main

import (
	"encoding/hex"
	"etero/evm"
	"os"
)

func main() {
	data := os.Args[1]
	code, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}

	evm.Run(code)
}
