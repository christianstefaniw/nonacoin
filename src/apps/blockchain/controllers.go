package blockchain

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	publicKey, privateKey, _ := genKeys()
	myWalletAddress := publicKey

	pub2, _, _ := genKeys()
	otherAddress := pub2

	bc := createBlockchain()
	t := createTransaction(nil, myWalletAddress, 9000)
	//t2 := createTransaction(myWalletAddress, otherAddress, privateKey, 245)
	valid := bc.appendTransactions(t)
	if !valid {
		fmt.Println("not valid 1")
	}
	bc.minePendingTransactions()

	t3 := createTransaction(myWalletAddress, otherAddress, 4)
	t3.sign(privateKey, myWalletAddress)
	valid2 := bc.appendTransactions(t3)
	if !valid2 {
		fmt.Println("not valid 2")
	}
	bc.minePendingTransactions()

	fmt.Fprint(w, bc)
	fmt.Fprintf(w, "\n\nchain valid: %t\n", bc.isChainValid())
	fmt.Fprintf(w, "my banace: %f", bc.getBalanceOfWallet(myWalletAddress))
}
