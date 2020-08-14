package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块哈希值
	Hash          string //当前区块哈希值
	Data          string //区块数据
}

//获取区块哈希值
func calculateHash(b *Block) string {
	blockData := string(b.Index+b.Timestamp) + b.PrevBlockHash + b.Data
	sum256 := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(sum256[:])
}

//生成新的区块
func GenerateNewBlock(preBlock *Block, data string) *Block {
	newBlock := &Block{
		Index:         preBlock.Index + 1,
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: preBlock.Hash,
		Data:          data,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

//生成创世区块
func GenerateGenesisBlock() *Block {
	basicBlock := &Block{
		Index:     -1,
		Timestamp: time.Now().Unix(),
	}
	return GenerateNewBlock(basicBlock, "Genesis Block")
}
