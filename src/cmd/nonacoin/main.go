package main

import (
	"context"
	"fmt"
	"nonacoin/src/apps/peer2peer"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/helpers"
	"nonacoin/src/wallet"
)

func main() {
	helpers.LoadDotEnv()
	wlt := wallet.NewWallet()
	node := peer2peer.NewPeerNode("127.0.0.1:8080", wlt)
	node.StartServer()
	cc, client := node.ConnectToClient("1234", "127.0.0.1:8081")
	resp, _ := client.SyncChain(context.Background(), &peer2peerpb.SyncChainRequest{Peer: "test"})
	fmt.Println(resp.GetNodes())
	fmt.Println(cc.GetState())
}
