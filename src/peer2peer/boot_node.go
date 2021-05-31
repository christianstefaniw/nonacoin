package peer2peer

type RoutingArray []string

// holds addresses of all peers
type BootNode struct {
	server       *peer2PeerServer
	routingArray RoutingArray
}

func newBootNode(addr string) *BootNode {
	node := new(BootNode)
	node.routingArray = NewRoutingArray()
	node.server = newPeer2PeerServer(addr, node)
	return node
}

func (n *BootNode) StartServer() {
	go n.server.start()
}

func (n *BootNode) WhichNode() int {
	return bootNode
}

func NewRoutingArray() RoutingArray {
	return make(RoutingArray, 0)
}
