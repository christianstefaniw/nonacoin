package peer2peer

import (
	"nonacoin/src/account"
	"nonacoin/src/blockchain"
	"nonacoin/src/pos"
	"nonacoin/src/transactions"
)

const (
	transThreshold = 1
	peerNode       = iota
	bootNode
)

type Node interface {
	StartServer()
	WhichNode() int
}

type PeerNode struct {
	validator       bool
	server          *peer2PeerServer
	stake           *pos.Stake
	transactionPool *transactions.TransactionPool
	blockchain      *blockchain.Blockchain
}

func NewPeerNode(addr string, acc *account.Account) *PeerNode {
	new := new(PeerNode)
	new.server = newPeer2PeerServer(addr, new)
	new.transactionPool = transactions.NewPool(transThreshold)
	new.blockchain = blockchain.NewBlockchain(acc)
	return new
}

func EmptyPeerNode() *PeerNode {
	return new(PeerNode)
}

func (p *PeerNode) WhichNode() int {
	return peerNode
}

func (p *PeerNode) StartServer() {
	p.server.start()
}

func (p *PeerNode) ConnectToClient(addr IP) interface{} {
	return p.server.setupClient(addr)
}

func (p *PeerNode) IsValidator() bool {
	return p.validator
}
