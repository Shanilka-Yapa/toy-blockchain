package block

import (
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index        uint64                    `json:"index"`
	Timestamp    time.Time                 `json:"timestamp"`
	Transactions []transaction.Transaction `json:"transactions"`
	PreviousHash string                    `json:"previous_hash"`
	Nonce        uint64                    `json:"nonce"`
	Hash         string                    `json:"hash"`
}
