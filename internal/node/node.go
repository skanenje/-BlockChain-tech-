package node

import (
    "swapBlock/internal/blockchain"
    "swapBlock/internal/agreement"
)

type Node struct {
    ID         string
    Blockchain *blockchain.Blockchain
}

func NewNode(id string) *Node {
    return &Node{
        ID:         id,
        Blockchain: blockchain.NewBlockchain(),
    }
}

func (n *Node) AddAgreement(agreement agreement.SaleAgreement) {
    n.Blockchain.AddBlock(agreement)
}