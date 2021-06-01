package peer2peer

import (
	"context"
	"fmt"
	"nonacoin/src/account"
	"nonacoin/src/peer2peer/bootnodepb"
	"nonacoin/src/peer2peer/peer2peerpb"
	"nonacoin/src/wallet"
)

func (*PeerNode) SyncChain(ctx context.Context, request *peer2peerpb.SyncChainRequest) (*peer2peerpb.SyncChainResponse, error) {
	response := &peer2peerpb.SyncChainResponse{
		//Chain: "sync chain: not implemented",
		Chain: "unimplemented",
		Nodes: request.Peer + " okok",
	}

	return response, nil
}

func (b *BootNode) Bootstrap(ctx context.Context, request *bootnodepb.BootstrapRequest) (*bootnodepb.BootstrapResponse, error) {
	peerAddr := request.GetAddr()
	wlt := wallet.NewWallet()
	acc := account.NewAccount(wlt)
	newPeer := NewPeerNode(peerAddr, acc)
	b.bootstrap(newPeer)

	response := &bootnodepb.BootstrapResponse{
		Success: true,
	}

	fmt.Println(b.routingTable)

	return response, nil
}
