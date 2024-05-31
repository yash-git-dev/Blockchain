package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	var b = &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
	}

	b.DeriveHash()

	return b
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.blocks[len(chain.blocks)-1]
	b := CreateBlock(data, lastBlock.Hash)
	chain.blocks = append(chain.blocks, b)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockchain()

	chain.AddBlock("abc")
	chain.AddBlock("pqs")
	chain.AddBlock("xyz")

	for _, block := range chain.blocks {
		fmt.Printf("data in block:%x\n", block.Hash)
	}
}
