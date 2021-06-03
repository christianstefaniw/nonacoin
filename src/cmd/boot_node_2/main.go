package main

import (
	"nonacoin/src/helpers"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer"
)

func main() {
	helpers.LoadDotEnv()
	addr := nonacoin.BOOT_NODES_ADDR[1]
	bootNode := peer2peer.NewBootNode(addr)
	bootNode.SyncRouteTable(peer2peer.LoadRoutingTable(addr))
	bootNode.StartServer()
}
