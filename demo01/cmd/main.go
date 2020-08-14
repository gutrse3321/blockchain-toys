package main

import "demo/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("send 1 BTC to Tomonori")
	bc.SendData("send 1 EOS to Tomonori")
	bc.Print()
}
