package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //区块创建时间戳
	Data          []byte //区块包含的数据
	PrevBlockHash []byte //前一个区块的哈希值
	Hash          []byte //区块自身的哈希值，用于校验区块数据有效
	Nonce         int
}

//构造一个区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}

	//工作证明
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//构造创世纪区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//哈希加密，设置block的Hash字段值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
