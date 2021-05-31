package peer2peer

import (
	"log"
	"net"
	"nonacoin/src/peer2peer/peer2peerpb"

	"google.golang.org/grpc"
)

type IP string

type ConnectedPeersTable map[IP]peer2peerpb.PeerToPeerServiceClient

type peer2PeerServer struct {
	routingArray RoutingArray
	peers        ConnectedPeersTable
	addr         string
	node         Node
}

func newPeer2PeerServer(addr string, node Node) *peer2PeerServer {
	return &peer2PeerServer{
		node:         node,
		addr:         addr,
		peers:        NewConnectedPeersTable(),
		routingArray: NewRoutingArray(),
	}
}

func NewConnectedPeersTable() ConnectedPeersTable {
	return make(ConnectedPeersTable)
}

func (s *peer2PeerServer) start() {
	s.routingArray = s.FindPeers()
	go s.startListening()
}

func (s *peer2PeerServer) startListening() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	peer2peerpb.RegisterPeerToPeerServiceServer(grpcServer, s.node.(peer2peerpb.PeerToPeerServiceServer))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *peer2PeerServer) FindPeers() RoutingArray {
	// connect to closest peers
	// use geographical locations
	// if location to attempted connection is greater than `x`, do not connect
	return make(RoutingArray, 0)
}

func (s *peer2PeerServer) setupClient(ip IP, addr string) (*grpc.ClientConn, peer2peerpb.PeerToPeerServiceClient) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to conect to %s: %v", addr, err)
	}

	s.peers[ip] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return conn, s.peers[ip]
}
