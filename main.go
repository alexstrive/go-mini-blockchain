package main

import (
	"fmt"
	"strconv"
)

// Taken from
// https://jeiwan.net/posts/building-blockchain-in-go-part-1/
// https://jeiwan.net/posts/building-blockchain-in-go-part-2/
// TODO: Learn more about Hashcash
// TODO: Learn more about Proof of Work, Proof of Stake
// https://jeiwan.net/posts/building-blockchain-in-go-part-3/
// https://jeiwan.net/posts/building-blockchain-in-go-part-4/
// https://jeiwan.net/posts/building-blockchain-in-go-part-5/
// https://jeiwan.net/posts/building-blockchain-in-go-part-6/
// https://jeiwan.net/posts/building-blockchain-in-go-part-7/
func main() {
	blockchain := NewBlockchain()

	blockchain.AddBlock("Some information")
	blockchain.AddBlock("Another useful information")

	for _, block := range blockchain.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}
