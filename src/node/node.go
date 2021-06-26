package node

import "nonacoin/src/peer2peer/peer2peerpb"

type NodeType int

type Node interface {
	peer2peerpb.PeerToPeerServiceServer
	WhichNode() NodeType
}

const (
	PeerNode NodeType = iota
	BootNode
)
