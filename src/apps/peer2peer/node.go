package peer2peer

import (
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/blockchain"
	"nonacoin/src/pos"
	"nonacoin/src/transactions"
	"nonacoin/src/wallet"

	"google.golang.org/grpc"
)

type PeerNode struct {
	validator       bool
	server          *peer2PeerServer
	stake           *pos.Stake
	wallet          *wallet.Wallet
	transactionPool *transactions.TransactionPool
	blockchain      *blockchain.Blockchain
}

func NewPeerNode(addr string, wlt *wallet.Wallet) *PeerNode {
	new := new(PeerNode)
	new.wallet = wlt
	new.server = newPeer2PeerServer(addr, new)
	return new
}

func (p *PeerNode) StartServer() {
	p.server.start()
}

func (p *PeerNode) ConnectToClient(pubKey, addr string) (*grpc.ClientConn, peer2peerpb.PeerToPeerServiceClient) {
	return p.server.setupClient(pubKey, addr)
}

func (p *PeerNode) IsValidator() bool {
	return p.validator
}
