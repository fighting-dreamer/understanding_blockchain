package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const MINIMUM_DIFFICULTY = 3

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

type BlockChain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockChain() *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}

	return b
}

func (bc *BlockChain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, len(bc.transactionPool))
	for i := 0; i < len(bc.transactionPool); i++ {
		transactions[i] = NewTransaction(bc.transactionPool[i].senderBlockchainAddress, bc.transactionPool[i].recipientBlockchainAddress, bc.transactionPool[i].value)
	}

	return transactions
}

func (bc *BlockChain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zerosString := strings.Repeat("0", difficulty)
	guessblock := Block{nonce, previousHash, 0, transactions}
	guessHashStr := fmt.Sprintf("%x", guessblock.Hash())

	return guessHashStr[:difficulty] == zerosString
}

func (bc *BlockChain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, previousHash, transactions, MINIMUM_DIFFICULTY) {
		nonce += 1
	}

	return nonce
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	block := &Block{nonce: nonce, previousHash: previousHash, transactions: transactions}
	block.timestamp = time.Now().UnixNano()
	return block
}

func (b *Block) Print() {
	fmt.Println("timestamp: ", b.timestamp)
	fmt.Println("nonce: ", b.nonce)
	fmt.Println("previousHash: ", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	// fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (bc *BlockChain) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 25))
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}
