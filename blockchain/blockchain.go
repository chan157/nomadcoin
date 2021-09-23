package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockChain struct {
	blocks []*block
}

/* Singleton Pattern
Singleton Pattern is only one instance
There is only one instance in our application
*/

// Only inside of package we can access
var b *blockChain

// Initialize only once
var once sync.Once

func GetBlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func createBlock(data string) *block {
	newBlock := block{
		Data:     data,
		Hash:     "",
		PrevHash: getLastHash(),
	}
	newBlock.calculateHash()
	return &newBlock
}

func getLastHash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocks-1].Hash
}

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func (b *blockChain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b *blockChain) AllBlocks() []*block {
	return b.blocks
}
