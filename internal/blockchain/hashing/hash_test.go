package hashing

import (
	"testing"
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
)

func TestHashDeterminism(t *testing.T) {

	b := block.Block{
		Index:        1,
		Timestamp:    time.Unix(1000, 0),
		Nonce:        10,
		PreviousHash: "abc123",
	}

	hash1 := CalculateBlockHash(&b)
	hash2 := CalculateBlockHash(&b)

	if hash1 != hash2 {
		t.Errorf("expected same hash, got different values")
	}
}
