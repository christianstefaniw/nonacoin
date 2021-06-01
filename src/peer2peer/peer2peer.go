package peer2peer

import (
	"log"
	"net"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer/bootnodepb"
	"nonacoin/src/peer2peer/peer2peerpb"

	"google.golang.org/grpc"
)

type ConnectedPeersTable map[string]peer2peerpb.PeerToPeerServiceClient
type ConnectedBootNodePeersTable map[string]bootnodepb.BootNodeServiceClient

type peer2PeerServer struct {
	routingTable RoutingTable
	peers        ConnectedPeersTable
	addr         string
	node         *PeerNode
}

func IsBootstrapConn(addr string) bool {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR {
		if addr == bootAddr {
			return true
		}
	}
	return false
}

func newPeer2PeerServer(addr string, node *PeerNode) *peer2PeerServer {
	return &peer2PeerServer{
		node:         node,
		addr:         addr,
		peers:        newConnectedPeersTable(),
		routingTable: newRoutingTable(),
	}
}

func newConnectedPeersTable() ConnectedPeersTable {
	return make(ConnectedPeersTable)
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

func (s *peer2PeerServer) FindPeers() RoutingTable {
	// connect to closest peers
	// use geographical locations
	// if location to attempted connection is greater than `x`, do not connect
	return make(RoutingTable)
}

func (s *peer2PeerServer) setupClient(addr string) (peer2peerpb.PeerToPeerServiceClient, error) {
	conn, err := DialClient(addr)
	if err != nil {
		return nil, err
	}

	s.peers[addr] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return s.peers[addr], nil
}

type peer2PeerBootNodeServer struct {
	routingTable RoutingTable
	peers        ConnectedBootNodePeersTable
	addr         string
	node         *BootNode
}

func newBootNodeServer(addr string, node *BootNode) *peer2PeerBootNodeServer {
	return &peer2PeerBootNodeServer{
		routingTable: newRoutingTable(),
		peers:        newConnectedBootNodePeersTable(),
		addr:         addr,
		node:         node,
	}
}

func newConnectedBootNodePeersTable() ConnectedBootNodePeersTable {
	return make(ConnectedBootNodePeersTable)
}

func (b *peer2PeerBootNodeServer) start() {
	b.connectToBootNodes()
	b.startListening()
}

func (b *peer2PeerBootNodeServer) startListening() {
	lis, err := net.Listen("tcp", b.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	bootnodepb.RegisterBootNodeServiceServer(grpcServer, b.node)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (b *peer2PeerBootNodeServer) connectToBootNodes() error {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR {
		conn, err := DialClient(bootAddr)
		if err != nil {
			return err
		}

		b.peers[bootAddr] = bootnodepb.NewBootNodeServiceClient(conn)
	}

	return nil
}
