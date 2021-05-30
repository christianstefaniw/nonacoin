package peer2peer

import (
	"context"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/transactions"
)

type peer2PeerNetwork struct {
	nodes           []*peerNode
	transactionPool transactions.TransactionPool
}

func (p *peer2PeerNetwork) AppendNode(n *peerNode) {
	p.nodes = append(p.nodes, n)
}

func newPeer2PeerServer() *peer2PeerNetwork {
	new := new(peer2PeerNetwork)
	new.nodes = make([]*peerNode, 0)
	return new
}

func (*peer2PeerNetwork) SyncChain(ctx context.Context, request *peer2peerpb.SyncChainRequest) (*peer2peerpb.SyncChainResponse, error) {
	response := &peer2peerpb.SyncChainResponse{
		//Chain: "sync chain: not implemented",
		Chain: "unimplemented",
		Nodes: request.Peer,
	}

	return response, nil
}
