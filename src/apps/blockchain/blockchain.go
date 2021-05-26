package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type transaction struct {
	toAddress   string
	fromAddress string
	time        time.Time
	hash        string
	amount      float64
	signature   string
}

type block struct {
	hash         string
	nonse        int64
	time         time.Time
	index        int
	transactions []*transaction
	prevHash     string
}

type blockchain struct {
	difficulty          uint8
	pendingTransactions []*transaction
	chain               []*block
}

func createBlockchain() *blockchain {
	bc := new(blockchain)
	bc.chain = make([]*block, 0)
	bc.pendingTransactions = make([]*transaction, 0)

	bc.chain = append(bc.chain, bc.createGenesisBlock())
	bc.difficulty = 2
	return bc
}

func (bc *blockchain) length() int {
	return len(bc.chain)
}

func (bc *blockchain) createGenesisBlock() *block {
	genesis := createBlock(nil, "", 0)
	return genesis
}

func (bc *blockchain) minePendingTransactions() {
	newBlock := createBlock(bc.pendingTransactions, bc.getLatestBlock().hash, bc.length())
	newBlock.mine(bc.difficulty)
	bc.appendBlock(newBlock)

	bc.pendingTransactions = nil
}

func (bc *blockchain) appendTransactions(trans ...*transaction) {
	bc.pendingTransactions = append(bc.pendingTransactions, trans...)
}

func (bc *blockchain) appendBlock(b *block) {
	bc.chain = append(bc.chain, b)
}

func (bc *blockchain) String() string {
	var output strings.Builder

	for _, b := range bc.chain {
		output.WriteString(fmt.Sprint(b))
	}

	output.WriteString(fmt.Sprintf("length: %d", bc.length()))

	return output.String()
}

func (bc *blockchain) isChainValid() bool {
	for _, blck := range bc.chain[1:] {
		currBlock := blck

		if !currBlock.hasValidTransactions() {
			return false
		}

		if currBlock.hash != currBlock.calculateHash() {
			return false
		}
	}

	return true
}

func (bc *blockchain) getLatestBlock() *block {
	return bc.chain[bc.length()-1]
}

func createBlock(transactions []*transaction, prevHash string, index int) *block {
	b := new(block)
	b.index = index
	b.time = time.Now()
	b.nonse = 0
	b.prevHash = prevHash
	b.transactions = transactions
	b.hash = b.calculateHash()
	return b
}

func (b *block) hasValidTransactions() bool {
	for _, trans := range b.transactions {
		if !trans.isValid() {
			return false
		}
	}
	return true
}

func (b *block) calculateHash() string {
	var hashTransactions strings.Builder
	var hashString string
	var hash [32]byte

	for _, t := range b.transactions {
		hashTransactions.WriteString(string(t.hash))
	}

	hashString = b.time.Format(time.ANSIC) + hashTransactions.String() + b.prevHash +
		strconv.Itoa(int(b.index)) + strconv.Itoa(int(b.nonse))
	hash = sha256.Sum256([]byte(hashString))
	return string(hash[:])
}

func (b *block) mine(difficulty uint8) bool {
	var target strings.Builder

	for i := 0; i < int(difficulty); i++ {
		target.WriteString(strconv.Itoa(i))
	}

	for b.hash[0:difficulty] != target.String() {
		b.nonse += 1
		b.hash = b.calculateHash()
	}

	return true
}

func (b *block) String() string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf("BLOCK %d \nhash: %x\ntime: %s\nprev hash: %x\nindex: %d\nnonse: %d\n",
		b.index, b.hash, b.time.Format(time.ANSIC), b.prevHash, b.index, b.nonse))

	output.WriteString("transactions: [\n")

	for _, t := range b.transactions {
		output.WriteString(fmt.Sprint(t))
	}

	output.WriteString("]\n\n")

	return output.String()
}

func createTransaction(fromAddress, toAddress string, privateKey *rsa.PrivateKey, amount float64) *transaction {
	t := new(transaction)
	t.toAddress = toAddress
	t.fromAddress = fromAddress
	t.time = time.Now()
	t.amount = amount
	t.hash = t.calculateHash()
	t.sign(privateKey, fromAddress)
	return t
}

func (t *transaction) calculateHash() string {
	var hashString string
	var hash [32]byte

	hashString = t.toAddress + t.fromAddress + t.time.Format(time.ANSIC) + fmt.Sprintf("%f", t.amount)
	hash = sha256.Sum256([]byte(hashString))
	return string(hash[:])
}

func (t *transaction) String() string {
	return fmt.Sprintf("\tto address: %s\n\tfrom address: %s\n\ttimestamp: %s\n\thash: %x\n\tamount: %f\n\n",
		t.toAddress, t.fromAddress, t.time.Format(time.ANSIC), t.hash, t.amount)
}

func (t *transaction) sign(privateKey *rsa.PrivateKey, walletAddress string) error {
	if t.fromAddress != walletAddress {
		return errors.New("cannot sign transaction for other wallet")
	}

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, []byte(t.hash), nil)
	if err != nil {
		return err
	}
	t.signature = string(signature)

	return nil
}

func (t *transaction) isValid() bool {
	if t.fromAddress == "" {
		return true
	}

	if t.signature == "" {
		return false
	}

	pubKey := t.pubKeyFromPem()
	err := rsa.VerifyPSS(pubKey, crypto.SHA256, []byte(t.calculateHash()), []byte(t.signature), nil)

	return err == nil
}

func (t *transaction) pubKeyFromPem() *rsa.PublicKey {
	pemBlock, _ := pem.Decode([]byte(t.fromAddress))

	parsedKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil
	}

	return parsedKey.(*rsa.PublicKey)
}
