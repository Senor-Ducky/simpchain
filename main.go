package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//This is a Block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

//This is a chain
type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	//this will join the previous block info with current block/next block.

	hash := sha256.Sum256(info)
	//hashing algorithm

	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

//first block or genesis block idk sounds cool
func Geneseis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Geneseis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}
