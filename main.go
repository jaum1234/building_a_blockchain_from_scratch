package main

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

func main() {

}
