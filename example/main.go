package main

import (
	"encoding/hex"
	"etero2/evm"
	"log"
	"os"
)

func main() {
	data := os.Args[1]
	code, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}

	err = evm.Run(code)
	if err != nil {
		log.Fatalf("failed to run contract: %v", err)
	}
}
