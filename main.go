package main

import (
	"fmt"
	blockchain "go-blockchain/src"
	"time"
)

func main() {
	blockchain := blockchain.InitBlockchain()
	start := time.Now()
	for i := 1; i <= 5000; i++ {
		blockchain.AddBlock([]byte(fmt.Sprintf("I'm %dº block", i)))
	}
	elapsed := time.Since(start)
	blockchain.ShowChain()

	fmt.Printf("Time taken to create blocks: %s\n", elapsed)
}
