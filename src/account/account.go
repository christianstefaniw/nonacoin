package account

import "nonacoin/src/wallet"

// manages transactions
type Account struct {
	wlt *wallet.Wallet
}

func NewAccount(wlt *wallet.Wallet) *Account {
	newAccount := new(Account)
	newAccount.wlt = wlt
	return newAccount
}

func (a *Account) GetWallet() *wallet.Wallet {
	return a.wlt
}
