package peer2peer

import (
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/wallet"
)

type PeerNode struct {
	Wallet *wallet.Wallet                      `json:"wallet"`
	Client peer2peerpb.PeerToPeerServiceClient `json:"client"`
}

func NewPeerNode(w *wallet.Wallet, c peer2peerpb.PeerToPeerServiceClient) *PeerNode {
	return &PeerNode{
		Wallet: w,
		Client: c,
	}
}
