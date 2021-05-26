package blockchain

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	privateKey, _ := genPrivateKey()
	publicKey := privateKey.PublicKey

	myWalletAddress, err := genPublicPem(&publicKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bc := createBlockchain()
	t := createTransaction(myWalletAddress, "you", privateKey, 20)
	t2 := createTransaction(myWalletAddress, "someone", privateKey, 245)
	valid := bc.appendTransactions(t, t2)
	if !valid {
		fmt.Println("not valid 1")
	}
	bc.minePendingTransactions()

	t3 := createTransaction(myWalletAddress, "person", privateKey, 1000000)
	valid2 := bc.appendTransactions(t3)
	if !valid2 {
		fmt.Println("not valid 2")
	}
	bc.minePendingTransactions()

	fmt.Fprint(w, bc)
	fmt.Fprintf(w, "\n\nchain valid: %t\n", bc.isChainValid())
	fmt.Fprintf(w, "my banace: %f", bc.getBalanceOfWallet(myWalletAddress))
}
