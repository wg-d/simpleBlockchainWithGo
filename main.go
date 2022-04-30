package main

import (
	"fmt"
	"time"
)

type Block struct {
	timestamp    time.Time
	transactions []string
	prevHash     []byte
	Hash         []byte
}

func main() {
	genesisTransactions := []string{"Izzy sent Will 50 bitcoin", "Will sent Izzy 30 bitcoin"}
	genesisBlock := NewBlock(genesisTransactions, []byte{})
	fmt.Println("--- First Block ---")
	printBlockInformation(genesisBlock)

	block2Transactions := []string{"John sent Izzy 30 bitcoin"}
	block2 := NewBlock(block2Transactions, genesisBlock.Hash)
	fmt.Println("--- Second Block ---")
	printBlockInformation(block2)

	block3Transactions := []string{"Will sent Izzy 45 bitcoin", "Izzy sent Will 10 bitcoin"}
	block3 := NewBlock(block3Transactions, block2.Hash)
	fmt.Println("--- Third Block ---")
	printBlockInformation(block3)
}
