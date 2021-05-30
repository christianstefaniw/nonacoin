package peer2peer

import (
	"log"
	"net"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"nonacoin/src/pos"
	"nonacoin/src/transactions"
	"nonacoin/src/wallet"

	"google.golang.org/grpc"
)

type PeerNode struct {
	peers           map[string]peer2peerpb.PeerToPeerServiceClient
	addr            string
	validator       bool
	stake           *pos.Stake
	wallet          *wallet.Wallet
	transactionPool *transactions.TransactionPool
}

func NewPeerNode(addr string, wlt *wallet.Wallet) *PeerNode {
	new := new(PeerNode)
	new.peers = make(map[string]peer2peerpb.PeerToPeerServiceClient)
	new.addr = addr
	new.wallet = wlt
	return new
}

func (p *PeerNode) IsValidator() bool {
	return p.validator
}

func (p *PeerNode) Start() {
	go p.startListening()
}

func (p *PeerNode) startListening() {
	lis, err := net.Listen("tcp", p.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	peer2peerpb.RegisterPeerToPeerServiceServer(grpcServer, p)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (p *PeerNode) SetupClient(pubKey, addr string) (*grpc.ClientConn, peer2peerpb.PeerToPeerServiceClient) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to conect to %s: %v", addr, err)
	}

	p.peers[pubKey] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return conn, p.peers[pubKey]
}
