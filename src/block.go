package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Block struct {
	hash         [32]byte
	previousHash [32]byte
	date         int64
	data         []byte
}

func NewBlock(data []byte, previousHash [32]byte) *Block {
	return &Block{
		hash:         [32]byte{},
		previousHash: previousHash,
		date:         time.Now().Unix(),
		data:         data,
	}
}

func (block *Block) loadHash() {
	block.hash = block.generateHash()
}

func (block *Block) generateHash() [32]byte {
	return sha256.Sum256(block.raw())
}

func (block *Block) raw() (raw []byte) {
	// raw = append(block.data, block.previousHash[:]...)
	// raw = append(raw, block.date2bytes()[:]...)
	raw = append(block.previousHash[:], block.date2bytes()...)
	raw = append(raw, block.data...)
	return raw
}

func (block *Block) date2bytes() []byte {
	date := make([]byte, 8)
	binary.LittleEndian.PutUint64(date, uint64(block.date))
	return date
}
