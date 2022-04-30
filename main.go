package main

import (
	"crypto/sha256"
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

func NewBlock(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevHash:     prevHash,
		Hash:         NewHash(currentTime, transactions, prevHash),
	}
}

func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}
