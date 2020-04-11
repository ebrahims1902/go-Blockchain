package main

import (
	"fmt"
	"strconv"

	"github.com/ebrahims1902/go-Blockchain/blockchain"
)

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash:%x\n", block.PrevHash)
		fmt.Printf("Data in Block:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
