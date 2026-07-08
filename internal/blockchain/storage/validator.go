package storage

import (
	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/validation"
)

// ValidateLoadedBlockchain verifies stored blockchain integrity.
func ValidateLoadedBlockchain(
	bc *chain.Blockchain,
	difficulty int,
) error {

	return validation.ValidateChain(bc, difficulty)
}
