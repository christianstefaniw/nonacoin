package peer2peer

import (
	"context"
	"math/rand"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer/peer2peerpb"
	"time"
)

// this is the first node on the network
// it implements functionality such as bootstrapping
// whever a new node wants to join the network, it must call the
// bootstrapping functionality
type BootNode struct {
	routingTable RoutingTable
	*PeerNode
}

func NewBootNode(addr string) *BootNode {
	node := new(BootNode)
	node.PeerNode = &PeerNode{}
	node.routingTable = emptyRoutingTable()
	node.server = newPeer2PeerServer(addr, node)
	node.connectToBootNodes()
	return node
}

func ConnectToBootNode() (peer2peerpb.BootNodeServiceClient, error) {
	rand.Seed(time.Now().UnixNano())
	conn, err := DialClient(nonacoin.BOOT_NODES_ADDR[rand.Intn(2)])
	if err != nil {
		return nil, err
	}
	return peer2peerpb.NewBootNodeServiceClient(conn), nil
}

func (b *BootNode) bootstrap(addr string) RoutingTable {
	ok := b.routingTable.IsActive(addr)
	if ok {
		return b.routingTable
	}

	b.routingTable.Add(addr)
	b.propagateConnection(addr)

	return b.routingTable
}

func (b *BootNode) propagateConnection(addr string) {
	request := &peer2peerpb.PropagateNewConnectionRequest{
		Addr: addr,
	}
	for _, bootPeer := range b.server.peers {
		bootPeer.(peer2peerpb.BootNodeServiceClient).PropagateNewConnection(context.Background(), request)
	}
}

// this method will connect boot peers together
func (b *BootNode) connectToBootNodes() error {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR {
		conn, err := DialClient(bootAddr)
		if err != nil {
			return err
		}

		b.server.peers[bootAddr] = peer2peerpb.NewBootNodeServiceClient(conn)
	}

	return nil
}

func (n *BootNode) StartServer() {
	n.server.start()
}

func (n *BootNode) WhichNode() NodeType {
	return bootNode
}
