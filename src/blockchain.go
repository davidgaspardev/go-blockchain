package blockchain

import "fmt"

type Blockchain struct {
	chain  []Block
	height int
}

func InitBlockchain() *Blockchain {
	blockchain := Blockchain{
		chain: []Block{},
	}

	blockchain.AddBlock([]byte("Genesis Block!"))

	return &blockchain
}

func (blockchain *Blockchain) AddBlock(data []byte) {
	if blockchain.height == 0 {
		genesisBlock := NewBlock(data, [32]byte{})
		genesisBlock.loadHash()
		blockchain.insertBlock(genesisBlock)
		return
	}

	previousHash := blockchain.chain[blockchain.height-1].hash
	block := NewBlock(data, previousHash)
	block.loadHash()
	blockchain.insertBlock(block)
}

func (blockchain *Blockchain) insertBlock(block *Block) {
	blockchain.chain = append(blockchain.chain, *block)
	blockchain.height++
}

func (blockchain *Blockchain) ValidateHash() bool {
	if blockchain.height > 1 {
		for i := 0; i < blockchain.height; i += 2 {
			if blockchain.chain[i].hash != blockchain.chain[i+2].previousHash {
				return false
			}
		}
	}
	return true
}

func (blockchain *Blockchain) ShowChain() {
	for i := 0; i < blockchain.height; i++ {
		block := &blockchain.chain[i]
		fmt.Printf("---> [ HEIGHT: %d]\n", i)
		fmt.Printf("hash: %x\n", block.hash)
		fmt.Printf("previousHash: %x\n", block.previousHash)
		fmt.Printf("date: %d\n", block.date)
		fmt.Printf("data: %s\n\n", block.data)
	}
}
