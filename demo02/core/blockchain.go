package core

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//添加一个区块到区块链
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, block)
}
