package peer2peer

import (
	"context"
	"fmt"
	"nonacoin/src/account"
	"nonacoin/src/peer2peer/bootnodepb"
	"nonacoin/src/peer2peer/peer2peerpb"
	"nonacoin/src/wallet"

	"google.golang.org/grpc"
)

func DialClient(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure())
}

func (*PeerNode) SyncChain(ctx context.Context, request *peer2peerpb.SyncChainRequest) (*peer2peerpb.SyncChainResponse, error) {
	response := &peer2peerpb.SyncChainResponse{
		//Chain: "sync chain: not implemented",
		Chain: "unimplemented",
		Nodes: request.Peer + " okok",
	}

	return response, nil
}

func (b *BootNode) Bootstrap(ctx context.Context, request *bootnodepb.BootstrapRequest) (*bootnodepb.BootstrapResponse, error) {
	peerAddr := request.GetAddr()
	wlt := wallet.NewWallet()
	acc := account.NewAccount(wlt)
	newPeer := newPeerNode(peerAddr, acc)
	b.bootstrap(newPeer)

	response := &bootnodepb.BootstrapResponse{
		Success: true,
	}

	fmt.Println(b.routingTable)
	fmt.Println(newPeer.server.routingTable)

	return response, nil
}

func (b *BootNode) RetrieveRoutingTable(ctx context.Context, request *bootnodepb.RetrieveRoutingTableRequest) (*bootnodepb.RetrieveRoutingTableResponse, error) {
	response := &bootnodepb.RetrieveRoutingTableResponse{
		Table: b.routingTable,
	}

	return response, nil
}

func (b *BootNode) PropagateNewConnection(ctx context.Context, request *bootnodepb.PropagateNewConnectionRequest) (*bootnodepb.PropagateNewConnectionResponse, error) {
	b.routingTable.Add(request.Addr)
	response := &bootnodepb.PropagateNewConnectionResponse{
		Success: b.routingTable.IsActive(request.Addr),
	}

	fmt.Println(b.routingTable, "table")

	return response, nil
}
