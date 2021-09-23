package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockChain struct {
	blocks []block
}

/* Singleton Pattern
Singleton Pattern is only one instance
There is only one instance in our application
*/

// Only inside of package we can access
var b *blockChain

func GetBlockChain() *blockChain {
	if b == nil {
		b = &blockChain{}
	}
	return b
}
