package ledger_test

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/ledger"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

func TestRejectOverspendingTransaction(t *testing.T) {

	l := ledger.NewLedger()

	// Give Alice initial funds
	faucetTx, err := transaction.NewCoinbase("Alice", 100)

	if err != nil {
		t.Fatalf("failed creating faucet transaction: %v", err)
	}

	err = l.ApplyTransaction(faucetTx)

	if err != nil {
		t.Fatalf("failed applying faucet transaction: %v", err)
	}

	// Alice tries to send more than balance
	tx, err := transaction.New(
		"Alice",
		"Bob",
		150,
	)

	if err != nil {
		t.Fatalf("failed creating transaction: %v", err)
	}

	err = l.ApplyTransaction(tx)

	if err == nil {
		t.Error("expected insufficient balance error")
	}

	if balance := l.GetBalance("Alice"); balance != 100 {
		t.Errorf(
			"balance changed after rejected transaction, got %f",
			balance,
		)
	}
}
