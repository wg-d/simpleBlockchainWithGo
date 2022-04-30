package main

import (
	"time"
)

type Block struct {
	timestamp    time.Time
	transactions []string
	prevHash     []byte
	Hash         []byte
}
