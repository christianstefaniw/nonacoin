package main

import (
	"nonacoin/src/helpers"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer"
)

func main() {
	helpers.LoadDotEnv()
	bootNode := peer2peer.NewBootNode(nonacoin.BOOT_NODES_ADDR[1])
	bootNode.StartServer()
}
