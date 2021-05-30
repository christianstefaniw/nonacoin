package peer2peer

import (
	"sync"
)

var _instance *peer2PeerNetwork
var _once sync.Once

func GetPeer2PeerInstance() *peer2PeerNetwork {
	_once.Do(func() {
		_instance = newPeer2PeerServer()
	})
	return _instance
}
