package blockchain

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    int
}

type Blockchain struct {
	Blocks []*Block
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

func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	b := CreateBlock(data, lastBlock.Hash)
	chain.Blocks = append(chain.Blocks, b)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
