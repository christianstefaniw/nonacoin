package main

import (
	"nonacoin/src/helpers"
	"nonacoin/src/net"
)

func main() {
	helpers.LoadDotEnv()
	net.ServeHTTP()
}
