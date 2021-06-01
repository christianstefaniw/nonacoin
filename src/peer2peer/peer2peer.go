package peer2peer

import (
	"context"
	"fmt"
	"log"
	"net"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer/peer2peerpb"

	"google.golang.org/grpc"
)

type ConnectedPeersTable map[string]peer2peerpb.PeerToPeerServiceClient

type RoutingTable map[string]bool

func (r RoutingTable) Add(addr string) bool {
	r[addr] = true
	return true
}

func (r RoutingTable) IsActive(addr string) bool {
	_, ok := r[addr]
	return ok
}

type Peer2PeerServer struct {
	routingTable RoutingTable
	peers        ConnectedPeersTable
	addr         string
	node         Node
}

func IsBootstrapConn(addr string) bool {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR {
		if addr == bootAddr {
			return true
		}
	}
	return false
}

func newPeer2PeerServer(addr string, node Node) *Peer2PeerServer {
	return &Peer2PeerServer{
		node:         node,
		addr:         addr,
		peers:        newConnectedPeersTable(),
		routingTable: newRoutingTable(),
	}
}

func newConnectedPeersTable() ConnectedPeersTable {
	return make(ConnectedPeersTable)
}

func emptyRoutingTable() RoutingTable {
	return make(RoutingTable)
}

func newRoutingTable() RoutingTable {
	client, err := ConnectToBootNode()
	if err != nil {
		log.Fatal(err)
	}
	routeTableMap, err := client.RetrieveRoutingTable(context.Background(), &peer2peerpb.RetrieveRoutingTableRequest{})
	if err != nil {
		return emptyRoutingTable()
	}
	fmt.Println(routeTableMap)
	rt := RoutingTable(routeTableMap.Table)
	return rt
}

func (s *Peer2PeerServer) start() {
	s.startListening()
}

func (s *Peer2PeerServer) startListening() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	switch s.node.WhichNode() {
	case bootNode:
		peer2peerpb.RegisterBootNodeServiceServer(grpcServer, s.node.(peer2peerpb.BootNodeServiceServer))
	case peerNode:
		peer2peerpb.RegisterPeerToPeerServiceServer(grpcServer, s.node.(peer2peerpb.PeerToPeerServiceServer))
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Peer2PeerServer) FindPeers() RoutingTable {
	// connect to closest peers
	// use geographical locations
	// if location to attempted connection is greater than `x`, do not connect
	return make(RoutingTable)
}

func (s *Peer2PeerServer) setupClient(addr string) (interface{}, error) {
	conn, err := DialClient(addr)
	if err != nil {
		return nil, err
	}

	s.peers[addr] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return s.peers[addr], nil
}
