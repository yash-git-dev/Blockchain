package blockchain

import (
	"log"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	b := CreateBlock(data, lastBlock.Hash)
	chain.Blocks = append(chain.Blocks, b)
}

func InitBlockchain() *Blockchain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	if err != nil {
		log.Printf(err.Error())
	}

	err = db.Update(func(txn *badger.Txn) error {
		if _, err = txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			genesis := Genesis()

			err = txn.Set(genesis.Hash, genesis.Serialize())
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

		}
		return err
	})

	return &Blockchain{lastHash, db}
}
