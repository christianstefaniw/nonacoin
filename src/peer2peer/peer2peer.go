package peer2peer

import (
	"log"
	"net"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer/bootnodepb"
	"nonacoin/src/peer2peer/peer2peerpb"

	"google.golang.org/grpc"
)

type IP string

type ConnectedPeersTable map[IP]peer2peerpb.PeerToPeerServiceClient

type NewPeer struct{}

type peer2PeerServer struct {
	routingTable RoutingTable
	peers        ConnectedPeersTable
	addr         IP
	node         Node
}

func (ip IP) String() string {
	return string(ip)
}

func (ip IP) IsBootstrapConn() bool {
	return ip.String() == nonacoin.BOOT_NODE_ADDR
}

func newPeer2PeerServer(addr string, node Node) *peer2PeerServer {
	return &peer2PeerServer{
		node:         node,
		addr:         IP(addr),
		peers:        NewConnectedPeersArray(),
		routingTable: NewRoutingTable(),
	}
}

func NewConnectedPeersArray() ConnectedPeersTable {
	return make(ConnectedPeersTable)
}

func (s *peer2PeerServer) start() {
	go s.startListening()
}

func (s *peer2PeerServer) startListening() {
	lis, err := net.Listen("tcp", s.addr.String())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	switch s.node.WhichNode() {
	case peerNode:
		peer2peerpb.RegisterPeerToPeerServiceServer(grpcServer, s.node.(peer2peerpb.PeerToPeerServiceServer))
	case bootNode:
		bootnodepb.RegisterBootNodeServiceServer(grpcServer, s.node.(bootnodepb.BootNodeServiceServer))
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *peer2PeerServer) FindPeers() RoutingTable {
	// connect to closest peers
	// use geographical locations
	// if location to attempted connection is greater than `x`, do not connect
	return make(RoutingTable)
}

func (s *peer2PeerServer) setupClient(addr IP) interface{} {
	conn, err := grpc.Dial(addr.String(), grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to conect to %s: %v", addr, err)
	}

	if addr.IsBootstrapConn() {
		return bootnodepb.NewBootNodeServiceClient(conn)
	}

	s.peers[addr] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return s.peers[addr]
}
