package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

//构造区块链
func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := &Blockchain{}
	blockchain.AppendBlock(genesisBlock)

	return blockchain
}

//添加一个新的区块
func (bc *Blockchain) AppendBlock(block *Block) {
	length := len(bc.Blocks)
	if length == 0 {
		bc.Blocks = append(bc.Blocks, block)
		return
	}

	if isValid(block, bc.Blocks[length-1]) {
		bc.Blocks = append(bc.Blocks, block)
	} else {
		log.Fatal("invalid block")
	}
}

//发送数据，创建一个新的区块，并加到区块链中
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	block := GenerateNewBlock(preBlock, data)
	bc.AppendBlock(block)
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d, Hash: %s\n", block.Index, block.Hash)
		fmt.Printf("Data: %s, Prev Hash: %s, Timestamp:%d\n\n", block.Data, block.PrevBlockHash, block.Timestamp)
	}
}

//判断两个区块old参数不是block参数的上一个节点
func isValid(block, old *Block) (valid bool) {
	if block.Index-1 != old.Index {
		return
	}

	if block.PrevBlockHash != old.Hash {
		return
	}

	if calculateHash(block) != block.Hash {
		return
	}

	return true
}
