package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockChain struct {
	blocks []block
}

func main() {
	chain := blockChain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")

	chain.listBlocks()
}

func (b *blockChain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}
func (b *blockChain) addBlock(data string) {
	newBlock := block{
		data:     data,
		hash:     "",
		prevHash: b.getLastHash(),
	}

	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	newBlock.hash = hexHash
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockChain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Hash : %s\n", block.hash)
		fmt.Printf("PrevHash : %s\n\n", block.prevHash)
	}
}
