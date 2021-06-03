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

func (r RoutingTable) Add(addr string) {
	r[addr] = true
}

func (r RoutingTable) IsActive(addr string) bool {
	_, ok := r[addr]
	return ok
}

func (r RoutingTable) ToMap() map[string]bool {
	return map[string]bool(r)
}

type Peer2PeerServer struct {
	routingTable RoutingTable
	peers        ConnectedPeersTable
	addr         string
	node         Node
}

func IsBootConn(addr string) bool {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR {
		if addr == bootAddr {
			return true
		}
	}
	return false
}

func newPeer2PeerServer(addr string, node Node) *Peer2PeerServer {
	return &Peer2PeerServer{
		node:  node,
		addr:  addr,
		peers: newConnectedPeersTable(),
	}
}

func newConnectedPeersTable() ConnectedPeersTable {
	return make(ConnectedPeersTable)
}

func emptyRoutingTable() RoutingTable {
	return make(RoutingTable)
}

func LoadRoutingTable(requestNodeAddr string) RoutingTable {
	client, err := ConnectToBootNode(requestNodeAddr)

	if err != nil {
		if err.Error() == "no boot nodes" {
			return emptyRoutingTable()
		} else {
			log.Fatal(err)
		}
	}

	routeTableMap, err := client.RetrieveRoutingTable(context.Background(), &peer2peerpb.RetrieveRoutingTableRequest{})
	if err != nil {
		return emptyRoutingTable()
	}

	fmt.Println(routeTableMap, "this was loaded from another boot peer")
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

func (s *Peer2PeerServer) setupPeerClientConn(addr string) (interface{}, error) {
	conn, err := DialClient(addr)
	fmt.Println(conn.GetState(), "2")
	if err != nil {
		return nil, err
	}

	s.peers[addr] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return s.peers[addr], nil
}
