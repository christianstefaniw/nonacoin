package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type transaction struct {
	sender   string
	reciever string
	time     time.Time
	hash     string
}

type block struct {
	hash         string
	nonse        uint16
	time         time.Time
	index        int64
	transactions []*transaction
	prevHash     string
}

type blockchain struct {
	pendingTransactions []*transaction
	blocks              []*block
}

func createBlock() *block {
	b := new(block)
	b.time = time.Now()
	b.index = 0
	b.nonse = 1234
	b.hash = b.calculateHash()
	return b
}

func (b *block) calculateHash() string {
	var hashTransactions bytes.Buffer
	var hashString string
	var hash [32]byte

	for _, t := range b.transactions {
		hashTransactions.WriteString(string(t.hash))
	}

	hashString = b.time.Format(time.ANSIC) + hashTransactions.String() + b.prevHash + strconv.Itoa(int(b.index)) + strconv.Itoa(int(b.nonse))
	fmt.Println(hashString)
	hash = sha256.Sum256([]byte(hashString))
	return string(hash[:])
}
