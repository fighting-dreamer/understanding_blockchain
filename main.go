package main

import "fmt"

func main() {
	bc := NewBlockChain()
	fmt.Println(bc)
	bc.CreateBlock(5, bc.LastBlock().Hash())
	bc.Print()
	bc.CreateBlock(10, bc.LastBlock().Hash())
	bc.Print()

	// b := &Block{nonce: 1}
	// fmt.Printf("%x\n", b.Hash())
}
