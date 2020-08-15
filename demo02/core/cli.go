package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
	命令行处理
*/

type CLI struct {
	Bc *Blockchain
}

//使用flag包解析命令行参数
func (c *CLI) Run() {
	c.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	//给addBlock添加 -data 标志
	addBlockData := addBlockCmd.String("data", "", "Block data")

	var err error
	switch os.Args[1] {
	case "addBlock":
		err = addBlockCmd.Parse(os.Args[2:])
	case "printChain":
		err = printChainCmd.Parse(os.Args[2:])
	default:
		c.printUsage()
		os.Exit(1)
	}
	if err != nil {
		log.Panic(err)
	}

	//检查用户提供的命令，解析相关的flag子命令
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		c.Bc.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}
}

func (c *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (c *CLI) validateArgs() {
	if len(os.Args) < 2 {

	}
}

//迭代打印区块链中的区块
func (c *CLI) printChain() {
	bci := c.Bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
