package blocks

import (
	"crypto/sha256"
	"fmt"
	trans "nonacoin/src/transactions"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	hash         string
	time         time.Time
	index        int
	transactions []*trans.Transaction
	prevHash     string
}

func CreateBlock(transactions []*trans.Transaction, prevHash string, index int) *Block {
	b := new(Block)
	b.index = index
	b.time = time.Now()
	b.prevHash = prevHash
	b.transactions = transactions
	b.hash = b.CalculateHash()
	return b
}

// check that all transactions in a block are valid
func (b *Block) HasValidTransactions() bool {
	for _, trans := range b.transactions {
		if !trans.IsValid() {
			return false
		}
	}
	return true
}

func (b *Block) CalculateHash() string {
	var hashTransactions strings.Builder
	var hashString string
	var hash [32]byte

	for _, t := range b.transactions {
		hashTransactions.WriteString(string(t.GetHash()))
	}

	hashString = b.time.Format(time.ANSIC) + hashTransactions.String() + b.prevHash +
		strconv.Itoa(int(b.index))
	hash = sha256.Sum256([]byte(hashString))
	return string(hash[:])
}

func (b *Block) String() string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf("BLOCK %d \nhash: %x\ntime: %s\nprev hash: %x\nindex: %d\n",
		b.index, b.hash, b.time.Format(time.ANSIC), b.prevHash, b.index))
	output.WriteString("transactions: [\n")

	for _, t := range b.transactions {
		output.WriteString(fmt.Sprint(t))
	}

	output.WriteString("]\n\n")

	return output.String()
}

func (b *Block) GetHash() string {
	return b.hash
}

func (b *Block) GetPrevHash() string {
	return b.prevHash
}

func (b *Block) GetTransactions() []*trans.Transaction {
	return b.transactions
}
