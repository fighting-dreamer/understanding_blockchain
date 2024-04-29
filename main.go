package main

func main() {
	bc := NewBlockChain()
	bc.Print()
	bc.AddTransaction("A", "B", 1.0)
	previousHash := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce, previousHash)
	bc.Print()
	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.001)
	previousHash = bc.LastBlock().Hash()
	nonce = bc.ProofOfWork()
	bc.CreateBlock(nonce, previousHash)
	bc.Print()

	// b := &Block{nonce: 1}
	// fmt.Printf("%x\n", b.Hash())
}
