package blockchain

import (
	"fmt"
	"net/http"
	"nonacoin/src/crypto"
	trans "nonacoin/src/transactions"
)

func test(w http.ResponseWriter, r *http.Request) {
	publicKey, privateKey, _ := crypto.GenKeys()
	myWalletAddress := publicKey

	pub2, _, _ := crypto.GenKeys()
	otherAddress := pub2

	bc := createBlockchain()
	t := trans.CreateTransaction(nil, myWalletAddress, 9000)
	valid := bc.appendTransactions(t)
	if !valid {
		fmt.Println("not valid 1")
	}
	bc.minePendingTransactions()

	t3 := trans.CreateTransaction(myWalletAddress, otherAddress, 4)
	t3.Sign(privateKey, myWalletAddress)
	valid2 := bc.appendTransactions(t3)
	if !valid2 {
		fmt.Println("not valid 2")
	}
	bc.minePendingTransactions()

	fmt.Fprint(w, bc)
	fmt.Fprintf(w, "\n\nchain valid: %t\n", bc.isChainValid())
	fmt.Fprintf(w, "my banace: %f\n\n\n", bc.getBalanceOfWallet(myWalletAddress))
	fmt.Fprintf(w, "transactions for %x:\n%s", myWalletAddress.String(), bc.getAllTransactionsForWallet(myWalletAddress))
}
