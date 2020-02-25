package main

import (
	"log"
	"sync"
	"time"

	"github.com/NavenduDuari/go-blockchain/bcgql"
	"github.com/NavenduDuari/go-blockchain/blockchain"
)

var mutex = &sync.Mutex{}

func main() {
	go func() {
		t := time.Now()
		genesisBlock := blockchain.Block{}
		genesisBlock = blockchain.Block{
			Index:     0,
			Timestamp: t.String(),
			Data:      0,
			Hash:      blockchain.CalculateHash(genesisBlock),
			PrevHash:  "",
		}

		mutex.Lock()
		blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)
		mutex.Unlock()
	}()

	log.Fatal(bcgql.Run())

}
