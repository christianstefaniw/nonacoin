package peer2peer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"nonacoin/src/helpers"
	"nonacoin/src/node"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer/peer2peerpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
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
	RoutingTable RoutingTable
	Peers        ConnectedPeersTable
	Addr         string
	Node         node.Node
}

func IsBootConn(addr string) bool {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR {
		if addr == bootAddr {
			return true
		}
	}
	return false
}

func NewPeer2PeerServer(addr string, node node.Node) *Peer2PeerServer {
	return &Peer2PeerServer{
		Node:  node,
		Addr:  addr,
		Peers: NewConnectedPeersTable(),
	}
}

func NewConnectedPeersTable() ConnectedPeersTable {
	return make(ConnectedPeersTable)
}

func EmptyRoutingTable() RoutingTable {
	return make(RoutingTable)
}

func ConnectToBootNode(exclude ...string) (peer2peerpb.BootNodeServiceClient, error) {
	var index int
	var currIter int
	var currBootAddr string
	var err error
	var conn *grpc.ClientConn

	uniqueRand := helpers.NewUniqueRand()
	numOfBootNodes := len(nonacoin.BOOT_NODES_ADDR)

	for {
		index = uniqueRand.Int(numOfBootNodes)
		currBootAddr = nonacoin.BOOT_NODES_ADDR[index]
		if currIter == numOfBootNodes-1 {
			return nil, errors.New("no boot nodes")
		}
		if helpers.SearchStringSlice(exclude, currBootAddr) {
			currIter++
			continue
		}

		conn, err = DialClient(currBootAddr, grpc.WithBlock())
		if err != nil {
			currIter++
			continue
		} else if conn.GetState() == connectivity.Ready {
			break
		}
	}

	if conn.GetState() != connectivity.Ready {
		return nil, err
	}

	return peer2peerpb.NewBootNodeServiceClient(conn), nil
}

func LoadRoutingTable(requestNodeAddr string) RoutingTable {
	client, err := ConnectToBootNode(requestNodeAddr)

	if err != nil {
		if err.Error() == "no boot nodes" {
			return EmptyRoutingTable()
		} else {
			log.Fatal(err)
		}
	}

	routeTableMap, err := client.RetrieveRoutingTable(context.Background(), &peer2peerpb.RetrieveRoutingTableRequest{})
	if err != nil {
		return EmptyRoutingTable()
	}

	fmt.Println(routeTableMap, "this was loaded from another boot peer")
	rt := RoutingTable(routeTableMap.Table)
	return rt
}

func (s *Peer2PeerServer) Start() {
	s.startListening()
}

func (s *Peer2PeerServer) startListening() {
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	switch s.Node.WhichNode() {
	case node.BootNode:
		peer2peerpb.RegisterBootNodeServiceServer(grpcServer, s.Node.(peer2peerpb.BootNodeServiceServer))
	case node.PeerNode:
		peer2peerpb.RegisterPeerToPeerServiceServer(grpcServer, s.Node.(peer2peerpb.PeerToPeerServiceServer))
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Peer2PeerServer) SetupPeerClientConn(addr string) (interface{}, error) {
	conn, err := DialClient(addr)
	fmt.Println(conn.GetState(), "2")
	if err != nil {
		return nil, err
	}

	s.Peers[addr] = peer2peerpb.NewPeerToPeerServiceClient(conn)

	return s.Peers[addr], nil
}
