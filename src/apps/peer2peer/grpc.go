package peer2peer

import (
	"context"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/pos"
)

type peer2PeerNetwork struct {
	validators []*pos.ValidatorNode
}

func (p *peer2PeerNetwork) AppendValidator(n *pos.ValidatorNode) {
	p.validators = append(GetPeer2PeerInstance().validators, n)
}

func newPeer2PeerServer() *peer2PeerNetwork {
	new := new(peer2PeerNetwork)
	new.validators = make([]*pos.ValidatorNode, 0)
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
