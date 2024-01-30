package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{[]*Block{NewBlock("Genesis Node", []byte{})}}
	return blockchain
}

func main() {
	blockchain := NewBlockchain()

	blockchain.AddBlock("Transaction[0x00|0x01|1.0]")
	blockchain.AddBlock("Transaction[0x01|0x00|1.0]")

	for _, block := range blockchain.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)
	}
}
