package main

import (
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

type BlockChain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock(0, "Init Hash")
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

func NewBlock(nonce int, previousHash string) *Block {
	block := &Block{nonce: nonce, previousHash: previousHash}
	block.timestamp = time.Now().UnixNano()
	return block
}

func (b *Block) Print() {
	fmt.Println("timestamp: ", b.timestamp)
	fmt.Println("nonce: ", b.nonce)
	fmt.Println("previousHash: ", b.previousHash)
	fmt.Println("transactions: ", b.transactions)
}

func (bc *BlockChain) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 25))
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}
