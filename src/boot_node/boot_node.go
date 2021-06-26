package bootnode

import (
	"context"
	"fmt"
	"math/rand"
	"nonacoin/src/node"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer"
	"nonacoin/src/peer2peer/peer2peerpb"
	peernode "nonacoin/src/peer_node"
	"time"
)

// this is the first node on the network
// it implements functionality such as bootstrapping
// whever a new node wants to join the network, it must call the
// bootstrapping functionality
type BootNode struct {
	routingTable peer2peer.RoutingTable
	*peernode.PeerNode
}

func NewBootNode(addr string) *BootNode {
	node := new(BootNode)
	node.PeerNode = &peernode.PeerNode{}
	node.routingTable = peer2peer.EmptyRoutingTable()
	node.Server = peer2peer.NewPeer2PeerServer(addr, node)
	node.connectToBootNodes()
	return node
}

func (b *BootNode) bootstrap(addr string) peer2peer.RoutingTable {
	rand.Seed(time.Now().UnixNano())
	newTable := selectRandAddresses(b.routingTable, len(b.routingTable))

	ok := b.routingTable.IsActive(addr)
	if ok {
		return b.routingTable
	}
	b.routingTable.Add(addr)
	b.propagateConnection(addr)

	return newTable
}

func selectRandAddresses(table peer2peer.RoutingTable, numRoutes int) peer2peer.RoutingTable {
	newTable := make(peer2peer.RoutingTable)
	for i := 0; i < numRoutes; i++ {
		for k := range table {
			newTable.Add(k)
			break
		}
	}
	return newTable
}

func (b *BootNode) propagateConnection(addr string) {
	request := &peer2peerpb.PropagateNewConnectionRequest{
		Addr: addr,
	}
	for _, bootPeer := range b.Server.Peers {
		bootPeer.(peer2peerpb.BootNodeServiceClient).PropagateNewConnection(context.Background(), request)
	}
}

// this method will connect boot peers together
func (b *BootNode) connectToBootNodes() error {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR[1:] {
		conn, err := peer2peer.DialClient(bootAddr)
		if err != nil {
			return err
		}

		b.Server.Peers[bootAddr] = peer2peerpb.NewBootNodeServiceClient(conn)
	}

	return nil
}

func (n *BootNode) StartServer() {
	n.Server.Start()
}

func (n *BootNode) WhichNode() node.NodeType {
	return node.BootNode
}

func (b *BootNode) Bootstrap(ctx context.Context, request *peer2peerpb.BootstrapRequest) (*peer2peerpb.BootstrapResponse, error) {
	rt := b.bootstrap(request.GetAddr())

	response := &peer2peerpb.BootstrapResponse{
		RoutingTable: rt.ToMap(),
	}

	fmt.Println(b.routingTable, "boot node routing table")
	fmt.Println(rt, "bootstrapped node routing table")

	return response, nil
}

func (b *BootNode) RetrieveRoutingTable(ctx context.Context, request *peer2peerpb.RetrieveRoutingTableRequest) (*peer2peerpb.RetrieveRoutingTableResponse, error) {
	response := &peer2peerpb.RetrieveRoutingTableResponse{
		Table: b.routingTable,
	}

	return response, nil
}

func (b *BootNode) PropagateNewConnection(ctx context.Context, request *peer2peerpb.PropagateNewConnectionRequest) (*peer2peerpb.PropagateNewConnectionResponse, error) {
	b.routingTable.Add(request.Addr)

	response := &peer2peerpb.PropagateNewConnectionResponse{
		Success: b.routingTable.IsActive(request.Addr),
	}

	fmt.Println(b.routingTable, "this was sent from another node")

	return response, nil
}
