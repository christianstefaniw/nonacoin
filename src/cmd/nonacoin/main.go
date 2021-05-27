package main

import (
	"nonacoin/src/grpc"
	"nonacoin/src/helpers"
)

func main() {
	helpers.LoadDotEnv()
	grpc.Serve()
}
