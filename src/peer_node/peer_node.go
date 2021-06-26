package peernode

import (
	"context"
	"nonacoin/src/account"
	"nonacoin/src/blockchain"
	"nonacoin/src/node"
	"nonacoin/src/peer2peer"
	"nonacoin/src/peer2peer/peer2peerpb"
	"nonacoin/src/transactions"
	"nonacoin/src/wallet"
)

type PeerNode struct {
	Server          *peer2peer.Peer2PeerServer
	TransactionPool *transactions.TransactionPool
	Blockchain      *blockchain.Blockchain
}

const (
	transThreshold = 1
)

func (p *PeerNode) SyncChain(ctx context.Context, request *peer2peerpb.SyncChainRequest) (*peer2peerpb.SyncChainResponse, error) {
	response := &peer2peerpb.SyncChainResponse{
		//Chain: "sync chain: not implemented",
		Chain: "unimplemented",
		Nodes: request.Peer + " okok",
	}

	return response, nil
}

func NewPeerNode(addr string) *PeerNode {
	wlt := wallet.NewWallet()
	acc := account.NewAccount(wlt)
	new := new(PeerNode)
	new.Server = peer2peer.NewPeer2PeerServer(addr, new)
	new.TransactionPool = transactions.NewPool(transThreshold)
	new.Blockchain = blockchain.NewBlockchain(acc)
	return new
}

func (p *PeerNode) SyncRouteTable(table peer2peer.RoutingTable) {
	p.Server.RoutingTable = table
}

func (p *PeerNode) GetAddr() string {
	return p.Server.Addr
}

func (p *PeerNode) WhichNode() node.NodeType {
	return node.PeerNode
}

func (p *PeerNode) StartServer() {
	go p.Server.Start()
}

func (p *PeerNode) ConnectToPeer(addr string) (interface{}, error) {
	return p.Server.SetupPeerClientConn(addr)
}
