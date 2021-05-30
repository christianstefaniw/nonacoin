package pos

import (
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/wallet"
)

type ValidatorNode struct {
	Wallet *wallet.Wallet                      `json:"wallet"`
	Client peer2peerpb.PeerToPeerServiceClient `json:"client"`
	Stake  int                                 `json:"stake"`
}
