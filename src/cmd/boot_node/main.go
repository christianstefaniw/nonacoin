package main

import (
	"nonacoin/src/helpers"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer"
)

func main() {
	helpers.LoadDotEnv()
	bootNode := peer2peer.NewBootNode(nonacoin.BOOT_NODE_ADDR)
	bootNode.StartServer()
}
