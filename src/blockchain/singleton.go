package blockchain

import "sync"

var _instance *Blockchain
var _once sync.Once

func GetBlockchainInstance() *Blockchain {
	_once.Do(func() {
		_instance = createBlockchain()
	})
	return _instance
}
