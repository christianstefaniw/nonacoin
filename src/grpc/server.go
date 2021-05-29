package grpc

import (
	"log"
	"net"
	"nonacoin/src/apps/peer2peer"
	"nonacoin/src/apps/peer2peer/peer2peerpb"
	"os"

	"google.golang.org/grpc"
)

func Serve() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	peer2peerpb.RegisterPeerToPeerServiceServer(grpcServer, peer2peer.GetPeer2PeerInstance())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
