package pos

import (
	"nonacoin/src/apps/peer2peer"
)

type Validators struct {
	validators []*ValidatorNode
}

type ValidatorNode struct {
	PeerNode *peer2peer.PeerNode `json:"peer"`
	Stake    *stake              `json:"stake"`
}

func NewValidatorNode(peer *peer2peer.PeerNode, s *stake) *ValidatorNode {
	return &ValidatorNode{
		PeerNode: peer,
		Stake:    s,
	}
}
