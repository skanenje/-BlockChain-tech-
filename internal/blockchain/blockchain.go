package blockchain

import (
	"sync"
	"swapBlock/internal/agreement"
)
type BlockChain struct{
	Chain []Block
	mutex sync.Mutex
}
func NewBlockChain()*BlockChain{
	return &BlockChain{
		Chain: []Block{createGenesisBlock()},
	}
}
func createGenesisBlock()Block{
	return *NewBlock(0, agreement.SaleAgreement{},"" )
}
func (b *BlockChain) Diffuclty()int{
	return 5 // abstract for now
}
func (b *BlockChain) AddBlock(){
	
}