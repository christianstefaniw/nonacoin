package main

import (
	"context"
	"fmt"
	"nonacoin/src/helpers"
	"nonacoin/src/peer2peer"
	"nonacoin/src/peer2peer/peer2peerpb"
)

func main() {
	helpers.LoadDotEnv()
	const addr = "127.0.0.1:8083"
	bootNode, _ := peer2peer.ConnectToBootNode()
	thisNode := peer2peer.NewPeerNode(addr)
	resp, _ := bootNode.Bootstrap(context.Background(), &peer2peerpb.BootstrapRequest{Addr: thisNode.GetAddr()})
	thisNode.SyncRouteTable(resp.GetRoutingTable())
	fmt.Println(resp.RoutingTable)
}