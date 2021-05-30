package main

import (
	"context"
	"encoding/json"
	"fmt"
	"nonacoin/src/apps/peer2peer"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/wallet"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, _ := grpc.Dial("localhost:8080", opts)
	defer cc.Close()

	client := peer2peerpb.NewPeerToPeerServiceClient(cc)
	wlt := wallet.NewWallet()

	peer := peer2peer.NewPeerNode(wlt, client)
	syncChain(client, peer)
}

func syncChain(client peer2peerpb.PeerToPeerServiceClient, peer *peer2peer.PeerNode) {
	marshaled, _ := json.Marshal(peer)

	request := &peer2peerpb.SyncChainRequest{
		Peer: string(marshaled),
	}
	resp, _ := client.SyncChain(context.Background(), request)

	fmt.Printf("%s\n", resp.Nodes)
}
