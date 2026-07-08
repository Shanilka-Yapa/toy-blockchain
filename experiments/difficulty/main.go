package main

import (
	"fmt"
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/mining"
)

func main() {

	difficulties := []int{1, 2, 3, 4}

	for _, difficulty := range difficulties {

		testBlock := block.Block{
			Index:        1,
			Timestamp:    time.Now(),
			Transactions: nil,
			PreviousHash: "0000000000000000",
		}

		result, err := mining.MineBlock(testBlock, difficulty)

		if err != nil {
			panic(err)
		}

		fmt.Println("Difficulty:", difficulty)
		fmt.Println("Attempts:", result.Attempts)
		fmt.Println("Duration:", result.Duration)
		fmt.Println("----------------------")
	}
}
