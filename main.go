package main

func main() {
	myBlockChainAddress := "my-blockchain_address"
	bc := NewBlockChain(myBlockChainAddress)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()
	bc.Print()

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.001)
	bc.Mining()
	bc.Print()

	// b := &Block{nonce: 1}
	// fmt.Printf("%x\n", b.Hash())
}
