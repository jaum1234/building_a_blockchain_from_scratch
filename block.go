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

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}

func (b *Block) HashIsEqual(p *Block) bool {

	for i, v := range b.PrevBlockHash {
		if v != p.Hash[i] {
			return false
		}
	}

	return true
}

func (b *Block) IsValid(p *Block) bool {
	if len(b.PrevBlockHash) != len(p.Hash) {
		return false
	}

	if !b.HashIsEqual(p) {
		return false
	}

	return true
}

/*
The first block in a blockchain is
usually called the genesis block.
*/
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
