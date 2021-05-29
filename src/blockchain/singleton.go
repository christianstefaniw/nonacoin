package blockchain

import "sync"

var _instance *blockchain
var _once sync.Once

func GetBlockchainInstance() *blockchain {
	_once.Do(func() {
		_instance = createBlockchain()
	})
	return _instance
}
