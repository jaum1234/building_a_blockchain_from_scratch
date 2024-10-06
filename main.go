package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/*
A block is a data structure data will contain valuable
information.

Different blockchains have varying block specifications.

In the bitcoin specification, Timestamp, PrevBlockHash
and Hash would be considered as header fields.

For the sake of simplicity, that project will consider a block
to be as simple as the data structure defined bellow, not following
any specific specification.
*/
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

/*
Calculating hashes is a costly operation, making the addition of new
blocks very difficult (what also makes there modification very
difficult), but also making blockchains much more secure.

For know, I'll simply concatenate the header fields to calculate the
hash.
*/
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) addBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := newBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}

func main() {

}
