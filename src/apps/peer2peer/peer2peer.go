package peer2peer

import (
	"fmt"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/wallet"
)

type Node struct {
	Wallet *wallet.Wallet                      `json:"wallet"`
	Client peer2peerpb.PeerToPeerServiceClient `json:"client"`
}

func (n *Node) String() string {
	out := fmt.Sprintf("wallet:\n%s", n.Wallet)
	return out
}

func NewPeer(w *wallet.Wallet, cc peer2peerpb.PeerToPeerServiceClient) *Node {
	return &Node{
		Wallet: w,
		Client: cc,
	}
}
