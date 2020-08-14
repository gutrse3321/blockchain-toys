package main

import (
	"demo/core"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Tomonori")
	bc.AddBlock("Send 2 BTC to Tomonori")

	for _, block := range bc.Blocks {
		fmt.Println("prev hash:", hex.EncodeToString(block.PrevBlockHash))
		fmt.Println("Data:", string(block.Data))
		fmt.Println("Hash:", hex.EncodeToString(block.Hash))

		pow := core.NewProofOfWork(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
