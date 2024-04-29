package main

import "fmt"

func main() {
	bc := NewBlockChain()
	fmt.Println(bc)
	bc.CreateBlock(5, "hash 1")
	bc.Print()
	bc.CreateBlock(10, "hash 2")
	bc.Print()
}
