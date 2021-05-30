package peer2peer

import (
	"log"
	"net"
	"nonacoin/src/apps/peer2peer/peer2peerpb"

	"google.golang.org/grpc"
)

type peer2PeerServer struct {
	peers map[string]peer2peerpb.PeerToPeerServiceClient
	addr  string
	node  *PeerNode
}

func newPeer2PeerServer(addr string, node *PeerNode) *peer2PeerServer {
	return &peer2PeerServer{
		node:  node,
		addr:  addr,
		peers: make(map[string]peer2peerpb.PeerToPeerServiceClient),
	}
}

func (s *peer2PeerServer) start() {
	go s.startListening()
}

func (s *peer2PeerServer) startListening() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	peer2peerpb.RegisterPeerToPeerServiceServer(grpcServer, s.node)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *peer2PeerServer) setupClient(pubKey, addr string) (*grpc.ClientConn, peer2peerpb.PeerToPeerServiceClient) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to conect to %s: %v", addr, err)
	}

	s.peers[pubKey] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return conn, s.peers[pubKey]
}
