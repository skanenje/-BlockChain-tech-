package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strings"
)

func calculateHash(block Block) string {
    record := fmt.Sprintf("%d%s%s%v%d", 
        block.Index, 
        block.Timestamp.String(), 
        block.PreviousHash, 
        block.Agreement, 
        block.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func (b *Block) Mine(difficulty int) {
    target := strings.Repeat("0", difficulty)
    for {
        b.Hash = calculateHash(*b)
        if strings.HasPrefix(b.Hash, target) {
            return
        }
        b.Nonce++
    }
}