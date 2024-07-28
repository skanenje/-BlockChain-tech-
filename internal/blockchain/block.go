package blockchain

import (
	"time"
	"swapBlock/internal/agreement"
)
type Block struct{
	Index int
	TimeStamp time.Time
	Agreement agreement.SaleAgreement
	Hash string
	PreviousHash string
	Nonce int
}
func NewBlock(index int, agreement agreement.SaleAgreement, prevHash string)*Block{
	return &Block{
		Index: index,
		TimeStamp: time.Now(),
		Agreement: agreement,
		PreviousHash: prevHash,
	}
}