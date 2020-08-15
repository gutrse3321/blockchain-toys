package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const (
	dbFile       = "blockchain.db"
	blocksBucket = "blocks"
)

type Blockchain struct {
	tip []byte
	Db  *bolt.DB
}

//构造区块链，打开数据库，若创世纪区块不存在，创建一个新的
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("没有找到blockchain，创建一个新的...")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	return &Blockchain{tip, db}
}

//添加一个区块到区块链
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	//获取最后一个块的哈希用于生成新的哈希
	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		return nil
	})

	bc.tip = newBlock.Hash

	return
}

//返回一个迭代器结构体引用，带块哈希(currentHash)和数据库的连接
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.Db}
}

//区块链迭代器，我们不想讲所有的快加载到内存中
//每当要对链中的快进行迭代式，就创建一个迭代器，里面存储了当前迭代的块哈希(currentHash)和数据库的连接
type BlockchainIterator struct {
	currentHash []byte
	Db          *bolt.DB
}

//迭代器的初始状态为链中的 tip，因此区块将从尾到头（创世块为头），也就是从最新的到最旧的进行获取。实际上，选择一个 tip 就是意味着给一条链“投票”
func (bi *BlockchainIterator) Next() *Block {
	var block *Block

	log.Println("iter times")

	err := bi.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(bi.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bi.currentHash = block.PrevBlockHash

	return block
}
