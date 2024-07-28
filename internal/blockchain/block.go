package blockchain

import (
    "time"
    "swapBlock/internal/agreement"
)

type Block struct {
    Index        int
    Timestamp    time.Time  // Add this field
    Agreement    agreement.SaleAgreement
    Hash         string
    PreviousHash string
    Nonce        int
}

func NewBlock(index int, agreement agreement.SaleAgreement, previousHash string) *Block {
    return &Block{
        Index:        index,
        Timestamp:    time.Now(),  // Set the timestamp
        Agreement:    agreement,
        PreviousHash: previousHash,
        Hash:         "",  // This will be set when we mine the block
        Nonce:        0,
    }
}