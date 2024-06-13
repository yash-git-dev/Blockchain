package main

import (
	"fmt"

	blockchain "github.com/yash-git-dev/go-blockchain/Blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("abc")
	chain.AddBlock("pqs")
	chain.AddBlock("xyz")

	for _, block := range chain.Blocks {

		fmt.Printf("data in block:%x\n", block.Hash)
		fmt.Printf("previous data in block:%x\n", block.PrevHash)

		pow := blockchain.NewProof(block)
		fmt.Println(pow.Validate())
	}
}
