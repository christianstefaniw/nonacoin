package peer2peer

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer/bootnodepb"
	"time"
)

type RoutingTable map[string]bool

func (r RoutingTable) Add(addr string) bool {
	r[addr] = true
	return true
}

func (r RoutingTable) IsActive(addr string) bool {
	_, ok := r[addr]
	return ok
}

// this is the first node on the network
// it implements functionality such as bootstrapping
// whever a new node wants to join the network, it must call the
// bootstrapping functionality
type BootNode struct {
	server       *peer2PeerBootNodeServer
	routingTable RoutingTable
}

func NewBootNode(addr string) *BootNode {
	node := new(BootNode)
	node.routingTable = newRoutingTable()
	node.server = newBootNodeServer(addr, node)
	return node
}

func ConnectToBootNode() (bootnodepb.BootNodeServiceClient, error) {
	rand.Seed(time.Now().UnixNano())
	conn, err := DialClient(nonacoin.BOOT_NODES_ADDR[rand.Intn(2)])
	if err != nil {
		return nil, err
	}
	return bootnodepb.NewBootNodeServiceClient(conn), nil
}

func (b *BootNode) bootstrap(p *PeerNode) RoutingTable {
	p.server.routingTable = b.routingTable

	_, ok := b.routingTable[p.server.addr]
	if ok {
		return b.routingTable
	}

	b.routingTable[p.server.addr] = true
	b.propagateConnection(p.server.addr)

	return b.routingTable
}

func (b *BootNode) propagateConnection(addr string) {
	request := &bootnodepb.PropagateNewConnectionRequest{
		Addr: addr,
	}
	for _, bootPeer := range b.server.peers {
		bootPeer.PropagateNewConnection(context.Background(), request)
	}
}

func (n *BootNode) StartServer() {
	n.server.start()
}

func (n *BootNode) WhichNode() NodeType {
	return bootNode
}

func newRoutingTable() RoutingTable {
	client, err := ConnectToBootNode()
	if err != nil {
		log.Fatal(err)
	}
	routeTableMap, err := client.RetrieveRoutingTable(context.Background(), &bootnodepb.RetrieveRoutingTableRequest{})
	if err != nil {
		return make(RoutingTable)
	}
	fmt.Println(routeTableMap)
	rt := RoutingTable(routeTableMap.Table)
	return rt
}
