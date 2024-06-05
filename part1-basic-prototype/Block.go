package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block define the block struct
type Block struct {
	// block height
	Height int64
	// time
	Timestamp int64
	// the transaction data
	Data []byte
	// the prev block hash
	PrevBlockHash []byte
	// the current block hash
	Hash []byte
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	// first, converts the block's timestamp to a byte array timestamp.
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// concatenates the previous block's hash b.PrevBlockHash, the current block's data b.Data, and the timestamp into a byte array headers
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// calculates the SHA-256 hash of headers using sha256.Sum256 and assigns the result to the current block's Hash field
	hash := sha256.Sum256(headers)
	// returns a pointer to the newly created block.
	b.Hash = hash[:]
}

// NewBlock is a function used to create a new block
// it takes two parameters:  the data of the current block and the hash of the previous block.
func NewBlock(data string, prevBlockHash []byte) *Block {
	// creates a new Block struct and initializes its fields
	block := &Block{
		1,                 // set block height
		time.Now().Unix(), // set to the current Unix timestamp
		[]byte(data),      // set to the provided data parameter
		prevBlockHash,     // set to the provided prevBlockHash parameter
		[]byte{},          // is initialized as empty
	}
	// calls block.SetHash() to calculate and set the hash value of the current block
	block.SetHash()

	// it returns a pointer to the newly created block
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	// the first block
	return NewBlock("Genesis Block", []byte{})
}
