package main_test

import (
	"os"
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/storage"
	"github.com/thulshani30/toy-blockchain/internal/config"
)

func TestBlockchainPersistenceWorkflow(t *testing.T) {

	path := "testdata/test-chain.json"

	defer os.RemoveAll("testdata")

	cfg := config.DefaultConfig()
	cfg.DataPath = path

	// Create new chain
	bc, err := storage.LoadBlockchain(cfg.DataPath)

	if err != nil {
		t.Fatalf("failed loading blockchain: %v", err)
	}

	// Save chain
	err = storage.SaveBlockchain(cfg.DataPath, bc)

	if err != nil {
		t.Fatalf("failed saving blockchain: %v", err)
	}

	// Reload chain
	loaded, err := storage.LoadBlockchain(cfg.DataPath)

	if err != nil {
		t.Fatalf("failed reloading blockchain: %v", err)
	}

	if len(loaded.Blocks) != 1 {
		t.Errorf(
			"expected genesis block after reload, got %d blocks",
			len(loaded.Blocks),
		)
	}
}
