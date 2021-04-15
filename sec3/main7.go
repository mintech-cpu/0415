package main

import "log"

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock() *Block {
	b := new(Block)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

}
