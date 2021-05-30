package account

// manages transactions
type Account struct {
	walletBalances map[string]float64
}

func createAccount() *Account {
	newAccount := new(Account)
	newAccount.walletBalances = make(map[string]float64)
	return newAccount
}
