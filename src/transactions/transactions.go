package transactions

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"nonacoin/src/crypto"
	"time"
)

type Transaction struct {
	toAddress   crypto.PublicKey
	fromAddress crypto.PublicKey
	time        time.Time
	hash        []byte
	amount      float64
	signature   []byte
}

func CreateTransaction(fromAddress, toAddress crypto.PublicKey, amount float64) *Transaction {
	t := new(Transaction)
	t.toAddress = toAddress
	t.fromAddress = fromAddress
	t.time = time.Now()
	t.amount = amount
	t.hash = t.calculateHash()
	return t
}

func (t *Transaction) calculateHash() []byte {
	var hashString string
	var hash [32]byte

	hashString = t.toAddress.String() + t.fromAddress.String() + t.time.Format(time.ANSIC) + fmt.Sprintf("%f", t.amount)
	hash = sha256.Sum256([]byte(hashString))
	return hash[:]
}

func (t *Transaction) String() string {
	return fmt.Sprintf("\tto address: %x\n\tfrom address: %x\n\ttimestamp: %s\n\thash: %x\n\tamount: %f\n\n",
		t.toAddress, t.fromAddress, t.time.Format(time.ANSIC), t.hash, t.amount)
}

// sign signs a transaction's hash
func (t *Transaction) Sign(privateKey crypto.PrivateKey, walletAddress crypto.PublicKey) error {
	if t.fromAddress.String() != walletAddress.String() {
		return errors.New("cannot sign transaction for other wallet")
	}

	signature := privateKey.Sign(t.hash)
	t.signature = signature

	return nil
}

func (t *Transaction) IsValid() bool {
	if t.fromAddress == nil {
		return true
	}

	if t.signature == nil || t.toAddress == nil {
		return false
	}

	if t.amount <= 0 {
		return false
	}

	return t.fromAddress.Verify(t.calculateHash(), t.signature)
}

func (t *Transaction) GetHash() []byte {
	return t.hash
}

func (t *Transaction) GetFromAddress() crypto.PublicKey {
	return t.fromAddress
}

func (t *Transaction) GetToAddress() crypto.PublicKey {
	return t.toAddress
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}
