package main

import (
	"context"
	"fmt"
	"nonacoin/src/helpers"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer"
	"nonacoin/src/peer2peer/bootnodepb"
)

func main() {
	helpers.LoadDotEnv()
	const addr = "127.0.0.1:8081"
	emptyPeer := peer2peer.EmptyPeerNode()
	bootNode, _ := emptyPeer.ConnectToClient(nonacoin.BOOT_NODE_ADDR).(bootnodepb.BootNodeServiceClient)

	resp, _ := bootNode.Bootstrap(context.Background(), &bootnodepb.BootstrapRequest{Addr: addr})
	fmt.Println(resp.Success)
}
