package main

import (
	"context"
	"fmt"
	"nonacoin/src/helpers"
	"nonacoin/src/peer2peer"
	"nonacoin/src/peer2peer/bootnodepb"
)

func main() {
	helpers.LoadDotEnv()
	const addr = "127.0.0.1:8082"
	bootNode, _ := peer2peer.ConnectToBootNode()
	resp, _ := bootNode.Bootstrap(context.Background(), &bootnodepb.BootstrapRequest{Addr: addr})
	fmt.Println(resp.Success)
}
