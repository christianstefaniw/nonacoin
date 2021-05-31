package peer2peer

import (
	"nonacoin/src/blockchain"
	"nonacoin/src/peer2peer/peer2peerpb"
	"nonacoin/src/pos"
	"nonacoin/src/transactions"
	"nonacoin/src/wallet"

	"google.golang.org/grpc"
)

const (
	transThreshold = 1
)

type Node interface {
	StartServer()
}

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
	new.transactionPool = transactions.NewPool(transThreshold)
	return new
}

func (p *PeerNode) StartServer() {
	p.server.start()
}

func (p *PeerNode) ConnectToClient(ip IP, addr string) (*grpc.ClientConn, peer2peerpb.PeerToPeerServiceClient) {
	return p.server.setupClient(ip, addr)
}

func (p *PeerNode) IsValidator() bool {
	return p.validator
}
