package main

import (
	bootnode "nonacoin/src/boot_node"
	"nonacoin/src/helpers"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer"
)

func main() {
	helpers.LoadDotEnv()
	addr := nonacoin.BOOT_NODES_ADDR[1]
	bootNode := bootnode.NewBootNode(addr)
	bootNode.SyncRouteTable(peer2peer.LoadRoutingTable(addr))
	bootNode.StartServer()
}
