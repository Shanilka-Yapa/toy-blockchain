package chain

//sync.RWMutex is used to ensure that the blockchain can be accessed by multiple goroutines safely. It allows multiple readers or one writer at a time
//Blocks is a slice of Block structs that represents the entire blockchain.

import (
	"sync"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
)

// Blockchain represents the complete blockchain.
type Blockchain struct {
	mu     sync.RWMutex
	Blocks []block.Block `json:"blocks"`
}