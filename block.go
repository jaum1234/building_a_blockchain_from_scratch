package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int32
	Difficulty    string
}

func (b *Block) ComputeHash(difficulty string) {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var nonce int32 = 0
	buf := new(bytes.Buffer)

	for {
		binary.Write(buf, binary.BigEndian, nonce)

		headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp, buf.Bytes(), []byte(difficulty)}, []byte{})
		hash := sha256.Sum256(headers)
		slice := hash[:]

		if strings.HasPrefix(hex.EncodeToString(slice), difficulty) {
			b.Hash = slice
			b.Nonce = nonce
			break
		} else {
			nonce++
			buf.Reset()
		}
	}
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	difficulty := "00"

	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0, difficulty}
	block.ComputeHash(difficulty)

	return block
}

func (b *Block) HashIsEqual(hash []byte) bool {
	for i, v := range b.Hash {
		if v != hash[i] {
			return false
		}
	}

	return true
}

func (b *Block) PreviousHashIsEqual(hash []byte) bool {

	for i, v := range b.PrevBlockHash {
		if v != hash[i] {
			return false
		}
	}

	return true
}

func (b *Block) IsValid(p *Block) bool {
	if len(b.PrevBlockHash) != len(p.Hash) {
		return false
	}

	if !p.HashIsEqual(b.PrevBlockHash) {
		return false
	}

	if p.Timestamp > b.Timestamp {
		return false
	}

	return true
}

func NewGenesisBlock() *Block {
	prev := "0000000000000000000000000000000000000000000000000000000000000000"

	return NewBlock("Genesis block", []byte(prev))
}
