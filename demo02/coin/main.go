package main

import (
	"demo/core"
)

func main() {
	bc := core.NewBlockchain()
	defer bc.Db.Close()

	cli := core.CLI{Bc: bc}
	cli.Run()
}
