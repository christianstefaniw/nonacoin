package blockchain

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	b := createBlock()
	fmt.Fprintf(w, "%x", b.hash)
}
