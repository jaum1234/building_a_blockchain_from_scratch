package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("I love my girlfriend.")
	bc.AddBlock("I love dogs.")
	bc.AddBlock("I love coffee.")

	//for _, block := range bc.blocks {
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	fmt.Println()
	//}

	fmt.Println(bc.blocks[2].IsValid(bc.blocks[1]))
}
