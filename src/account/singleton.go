package account

import "sync"

var _once sync.Once
var _instance *Account

func GetAccountInstance() *Account {
	_once.Do(func() {
		_instance = createAccount()
	})
	return _instance
}
