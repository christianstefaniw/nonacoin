package peer2peer

import (
	"nonacoin/src/account"
	"nonacoin/src/blockchain"
	"nonacoin/src/peer2peer/peer2peerpb"
	"nonacoin/src/transactions"
)

type NodeType int

type Node interface {
	peer2peerpb.PeerToPeerServiceServer
	WhichNode() NodeType
}

const (
	transThreshold          = 1
	peerNode       NodeType = iota
	bootNode
)

type PeerNode struct {
	server          *Peer2PeerServer
	transactionPool *transactions.TransactionPool
	blockchain      *blockchain.Blockchain
}

func newPeerNode(addr string, acc *account.Account) *PeerNode {
	new := new(PeerNode)
	new.server = newPeer2PeerServer(addr, new)
	new.transactionPool = transactions.NewPool(transThreshold)
	new.blockchain = blockchain.NewBlockchain(acc)
	return new
}

func (p *PeerNode) GetAddr() string {
	return p.server.addr
}

func (p *PeerNode) WhichNode() NodeType {
	return peerNode
}

func (p *PeerNode) StartServer() {
	go p.server.start()
}

func (p *PeerNode) ConnectToClient(addr string) (interface{}, error) {
	return p.server.setupClient(addr)
}
