package blockchain

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	bc := createBlockchain()
	t := createTransaction("me", "you", 20)
	t2 := createTransaction("bob", "joe", 502.1)
	t3 := createTransaction("john", "guy", 32)
	bc.appendTransaction(t)
	bc.appendTransaction(t2)
	bc.minePendingTransactions()
	bc.appendTransaction(t3)
	bc.minePendingTransactions()
	fmt.Fprint(w, bc)
}
