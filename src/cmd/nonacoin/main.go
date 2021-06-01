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
	const addr = "127.0.0.1:8082"
	bootNode, _ := peer2peer.ConnectToBootNode()
	resp, err := bootNode.Bootstrap(context.Background(), &peer2peerpb.BootstrapRequest{Addr: addr})
	fmt.Println(err)
	fmt.Println(resp.Success)
}
