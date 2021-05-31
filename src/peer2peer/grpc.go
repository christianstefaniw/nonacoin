package peer2peer

import (
	"context"
	"nonacoin/src/peer2peer/peer2peerpb"
)

func (*PeerNode) SyncChain(ctx context.Context, request *peer2peerpb.SyncChainRequest) (*peer2peerpb.SyncChainResponse, error) {
	response := &peer2peerpb.SyncChainResponse{
		//Chain: "sync chain: not implemented",
		Chain: "unimplemented",
		Nodes: request.Peer + " okok",
	}

	return response, nil
}

func (*BootNode) Bootstrap(ctx context.Context, request *peer2peerpb.BootstrapRequest) (*peer2peerpb.BootstrapResponse, error) {
	respone := &peer2peerpb.BootstrapResponse{
		RoutingArray: make([]string, 0),
	}
	respone.RoutingArray = append(respone.RoutingArray, "127.0.0.1")

	return respone, nil
}
