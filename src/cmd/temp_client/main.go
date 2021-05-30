package main

import (
	"nonacoin/src/apps/peer2peer"
	"nonacoin/src/helpers"
	"nonacoin/src/wallet"
)

func main() {
	helpers.LoadDotEnv()
	wlt := wallet.NewWallet()
	wlt.SetPubKey("1234")
	node := peer2peer.NewPeerNode("127.0.0.1:8081", wlt)
	node.Start()
}
