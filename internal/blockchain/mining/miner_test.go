package mining_test

import (
	"strings"
	"testing"
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/mining"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

func TestMiningMeetsDifficulty(t *testing.T) {

	testBlock := block.Block{
		Index:        1,
		Timestamp:    time.Now(),
		Transactions: []transaction.Transaction{},
		PreviousHash: "0000",
	}

	difficulty := 2

	result, err := mining.MineBlock(testBlock, difficulty)

	if err != nil {
		t.Fatalf("mining failed: %v", err)
	}

	if !strings.HasPrefix(result.Block.Hash, strings.Repeat("0", difficulty)) {
		t.Errorf(
			"hash %s does not satisfy difficulty",
			result.Block.Hash,
		)
	}

	if result.Attempts == 0 {
		t.Error("expected mining attempts greater than zero")
	}
}
