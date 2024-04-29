package main

func main() {
	bc := NewBlockChain()
	bc.Print()
	bc.AddTransaction("A", "B", 1.0)
	bc.CreateBlock(5, bc.LastBlock().Hash())
	bc.Print()
	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.001)
	bc.CreateBlock(10, bc.LastBlock().Hash())
	bc.Print()

	// b := &Block{nonce: 1}
	// fmt.Printf("%x\n", b.Hash())
}
