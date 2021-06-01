package peer2peer

import (
	"nonacoin/src/account"
	"nonacoin/src/blockchain"
	"nonacoin/src/peer2peer/peer2peerpb"
	"nonacoin/src/transactions"
)

type NodeType int

const (
	transThreshold          = 1
	peerNode       NodeType = iota
	bootNode
)

type PeerNode struct {
	server          *peer2PeerServer
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
	p.server.start()
}

func (p *PeerNode) ConnectToClient(addr string) (peer2peerpb.PeerToPeerServiceClient, error) {
	return p.server.setupClient(addr)
}
