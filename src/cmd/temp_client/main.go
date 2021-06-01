package main

import (
	"nonacoin/src/account"
	"nonacoin/src/helpers"
	"nonacoin/src/peer2peer"
	"nonacoin/src/wallet"
)

func main() {
	helpers.LoadDotEnv()
	wlt := wallet.NewWallet()
	acc := account.NewAccount(wlt)
	node := peer2peer.NewPeerNode("127.0.0.1:8081", acc)
	node.StartServer()
}
