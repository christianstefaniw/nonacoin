package peer2peer

import (
	"context"
	"fmt"
	"nonacoin/src/account"
	"nonacoin/src/peer2peer/peer2peerpb"
	"nonacoin/src/wallet"

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
	fmt.Println("ok")
	peerAddr := request.GetAddr()
	wlt := wallet.NewWallet()
	acc := account.NewAccount(wlt)
	newPeer := newPeerNode(peerAddr, acc)
	b.bootstrap(newPeer)

	response := &peer2peerpb.BootstrapResponse{
		Success: true,
	}

	fmt.Println(b.routingTable)
	fmt.Println(newPeer.server.routingTable)

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

	fmt.Println(b.routingTable, "table")

	return response, nil
}
