package chain_test

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/validation"
)

func TestGenesisBlockCreation(t *testing.T) {

	bc := chain.NewBlockchain()

	if len(bc.Blocks) != 1 {
		t.Fatalf("expected blockchain to contain genesis block")
	}

	genesis := bc.Blocks[0]

	if genesis.Index != 0 {
		t.Errorf("expected genesis index 0, got %d", genesis.Index)
	}

	if genesis.PreviousHash != chain.GenesisPreviousHash {
		t.Errorf("invalid genesis previous hash")
	}
}

func TestValidGenesisChain(t *testing.T) {

	bc := chain.NewBlockchain()

	err := validation.ValidateChain(bc, 3)

	if err != nil {
		t.Errorf("expected valid chain, got %v", err)
	}
}
