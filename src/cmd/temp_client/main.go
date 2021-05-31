package main

import (
	"nonacoin/src/helpers"
	"nonacoin/src/peer2peer"
	"nonacoin/src/wallet"
)

func main() {
	helpers.LoadDotEnv()
	wlt := wallet.NewWallet()
	wlt.SetPubKey("1234")
	node := peer2peer.NewPeerNode("127.0.0.1:8081", wlt)
	node.StartServer()
}
