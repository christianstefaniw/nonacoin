package peer2peer

type RoutingTable map[IP]bool

// this is the first node on the network
// it implements functionality such as bootstrapping
// whever a new node wants to join the network, it must call the
// bootstrapping functionality
type BootNode struct {
	server       *peer2PeerServer
	routingTable RoutingTable
}

func NewBootNode(addr string) *BootNode {
	node := new(BootNode)
	node.routingTable = NewRoutingTable()
	node.server = newPeer2PeerServer(addr, node)
	return node
}

func (b *BootNode) bootstrap(p *PeerNode) RoutingTable {
	_, ok := b.routingTable[p.server.addr]
	if ok {
		return b.routingTable
	}

	b.routingTable[p.server.addr] = true
	p.server.routingTable = b.routingTable

	return b.routingTable
}

func (n *BootNode) StartServer() {
	n.server.start()
}

func (n *BootNode) WhichNode() int {
	return bootNode
}

func NewRoutingTable() RoutingTable {
	return make(RoutingTable)
}
