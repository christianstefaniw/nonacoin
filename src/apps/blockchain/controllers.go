package blockchain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func test(w http.ResponseWriter, r *http.Request) {
	var privatePem strings.Builder
	var publicPem strings.Builder

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2024)
	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when dumping publickey: %s\n", err)
		return
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	pem.Encode(&publicPem, publicKeyBlock)
	pem.Encode(&privatePem, privateKeyBlock)

	myWalletAddress := publicPem.String()

	bc := createBlockchain()
	t := createTransaction(myWalletAddress, "you", privateKey, 20)
	t2 := createTransaction(myWalletAddress, "someone", privateKey, 245)
	bc.appendTransactions(t, t2)
	bc.minePendingTransactions()

	t3 := createTransaction(myWalletAddress, "person", privateKey, 1000000)
	bc.appendTransactions(t3)
	bc.minePendingTransactions()

	bc.chain[1].transactions[0].amount = 432

	fmt.Fprint(w, bc)
	fmt.Fprintf(w, "\n\nchain valid: %t\n", bc.isChainValid())
}
