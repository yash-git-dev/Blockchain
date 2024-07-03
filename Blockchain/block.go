package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	var b = &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}

	pow := NewProof(b)
	nonce, hash := pow.Run()

	b.Nonce = nonce
	b.Hash = hash[:]
	return b
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
