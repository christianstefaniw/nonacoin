package peer2peer

import (
	"context"
	"fmt"
	"nonacoin/src/peer2peer/peer2peerpb"

	"google.golang.org/grpc"
)

func DialClient(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure())
}

func (p *PeerNode) SyncChain(ctx context.Context, request *peer2peerpb.SyncChainRequest) (*peer2peerpb.SyncChainResponse, error) {
	response := &peer2peerpb.SyncChainResponse{
		//Chain: "sync chain: not implemented",
		Chain: "unimplemented",
		Nodes: request.Peer + " okok",
	}

	return response, nil
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
