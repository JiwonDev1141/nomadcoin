package blockchain

package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain 

func GetBlockChain()*blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}