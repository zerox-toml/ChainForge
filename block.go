package main

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "time"
)

// Block represents a single block in the blockchain
type Block struct {
    Index     int    // Position of the block in the chain
    Timestamp string // Time of block creation
    Data      string // Transaction or data stored in the block
    PrevHash  string // Hash of the previous block
    Hash      string // SHA-256 hash of the block
}

// NewBlock creates a new block using provided data and previous hash
func NewBlock(index int, data string, prevHash string) Block {
    timestamp := time.Now().Format(time.RFC3339)
    block := Block{
        Index:     index,
        Timestamp: timestamp,
        Data:      data,
        PrevHash:  prevHash,
    }
    block.Hash = CalculateHash(block)
    return block
}

// CalculateHash computes the SHA-256 hash of a block's contents
func CalculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
} 