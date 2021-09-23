package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data    string
	hash    string
	preHash string
}

func main() {
	genesisBlock := block{
		data:    "Genesis Block",
		hash:    "",
		preHash: "",
	}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.preHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}
