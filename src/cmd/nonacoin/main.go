package main

import (
	"context"
	"fmt"
	"log"
	"nonacoin/src/helpers"
	"nonacoin/src/peer2peer"
	"nonacoin/src/peer2peer/peer2peerpb"
)

func main() {
	helpers.LoadDotEnv()
	const addr = "127.0.0.1:8086"
	bootNode, err := peer2peer.ConnectToBootNode()
	if err != nil {
		log.Fatal(err)
	}
	thisNode := peer2peer.NewPeerNode(addr)
	resp, _ := bootNode.Bootstrap(context.Background(), &peer2peerpb.BootstrapRequest{Addr: thisNode.GetAddr()})
	thisNode.SyncRouteTable(resp.GetRoutingTable())
	fmt.Println(resp.RoutingTable)

}
