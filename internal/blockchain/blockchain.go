package blockchain

import (
    "sync"

    "swapBlock/internal/agreement"
)

type Blockchain struct {
    Chain  []Block
    mutex  sync.Mutex
}

func NewBlockchain() *Blockchain {
    return &Blockchain{
        Chain: []Block{createGenesisBlock()},
    }
}

func (bc *Blockchain) AddBlock(agreement agreement.SaleAgreement) {
    bc.mutex.Lock()
    defer bc.mutex.Unlock()

    prevBlock := bc.Chain[len(bc.Chain)-1]
    newBlock := NewBlock(len(bc.Chain), agreement, prevBlock.Hash)
    newBlock.Mine(bc.GetDifficulty())
    bc.Chain = append(bc.Chain, *newBlock)
}

func createGenesisBlock() Block {
    return *NewBlock(0, agreement.SaleAgreement{}, "")
}

func (bc *Blockchain) GetDifficulty() int {
    return 4 // This could be dynamic based on certain conditions
}