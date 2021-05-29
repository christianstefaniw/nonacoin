package peer2peer

import (
	"context"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
)

type peer2PeerServer struct {
	peers []*Node
}

func (p *peer2PeerServer) AppendPeer(n *Node) {
	p.peers = append(GetPeer2PeerInstance().peers, n)
}

func newPeer2PeerServer() *peer2PeerServer {
	new := new(peer2PeerServer)
	new.peers = make([]*Node, 0)
	return new
}

func (*peer2PeerServer) SyncChain(ctx context.Context, request *peer2peerpb.SyncChainRequest) (*peer2peerpb.SyncChainResponse, error) {
	response := &peer2peerpb.SyncChainResponse{
		//Chain: "sync chain: not implemented",
		Chain: "unimplemented",
		Nodes: request.Peer,
	}

	return response, nil
}
