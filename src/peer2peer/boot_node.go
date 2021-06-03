package peer2peer

import (
	"context"
	"errors"
	"math/rand"
	"nonacoin/src/helpers"
	"nonacoin/src/nonacoin"
	"nonacoin/src/peer2peer/peer2peerpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
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

func (b *BootNode) bootstrap(addr string) RoutingTable {
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

func selectRandAddresses(table RoutingTable, numRoutes int) RoutingTable {
	newTable := make(RoutingTable)
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
	for _, bootPeer := range b.server.peers {
		bootPeer.(peer2peerpb.BootNodeServiceClient).PropagateNewConnection(context.Background(), request)
	}
}

// this method will connect boot peers together
func (b *BootNode) connectToBootNodes() error {
	for _, bootAddr := range nonacoin.BOOT_NODES_ADDR[1:] {
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
