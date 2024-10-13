package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
)

type Blockchain struct {
	blocks []*Block
}

type BlockToBeAdded struct {
	Data string `json:"data"`
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) IsBroken() bool {
	for i, v := range bc.blocks {
		timestamp := []byte(strconv.FormatInt(v.Timestamp, 10))
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, v.Nonce)

		headers := bytes.Join([][]byte{v.PrevBlockHash, v.Data, timestamp, buf.Bytes(), []byte(v.Difficulty)}, []byte{})
		hash := sha256.Sum256(headers)
		slice := hash[:]

		if !v.HashIsEqual(slice) {
			return false
		}

		if !strings.HasPrefix(hex.EncodeToString(v.Hash), v.Difficulty) {
			return false
		}

		if i == 0 {
			if !v.PreviousHashIsEqual([]byte("0000000000000000000000000000000000000000000000000000000000000000")) {
				return false
			}
		} else {
			if !v.IsValid(bc.blocks[i-1]) {
				return false
			}
		}
	}

	return true
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
