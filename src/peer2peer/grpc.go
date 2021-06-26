package peer2peer

import (
	"google.golang.org/grpc"
)

func DialClient(addr string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	opts = append(opts, options...)
	return grpc.Dial(addr, opts...)

}
