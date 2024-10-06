package main

/*
In its essence, blockchain is just a public distributed
database of records.

A new block can only be added after the consent of
other keepers of the database.
*/

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
