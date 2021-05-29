package peer2peer

import (
	"sync"
)

var _instance *peer2PeerServer
var _once sync.Once

func GetPeer2PeerInstance() *peer2PeerServer {
	_once.Do(func() {
		_instance = newPeer2PeerServer()
	})
	return _instance
}
