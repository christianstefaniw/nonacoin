package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type transaction struct {
	sender   string
	receiver string
	time     time.Time
	hash     string
	amount   float64
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
	bc.difficulty = 3
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

func (bc *blockchain) appendTransaction(trans *transaction) {
	bc.pendingTransactions = append(bc.pendingTransactions, trans)
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
	output.WriteString(fmt.Sprintf("hash: %x\ntime: %s\nprev hash: %x\nindex: %d\nnonse: %d\n",
		b.hash, b.time.Format(time.ANSIC), b.prevHash, b.index, b.nonse))

	output.WriteString("transactions: [\n")

	for _, t := range b.transactions {
		output.WriteString(fmt.Sprint(t))
	}

	output.WriteString("]\n\n")

	return output.String()
}

func createTransaction(sender, receiver string, amount float64) *transaction {
	t := new(transaction)
	t.sender = sender
	t.receiver = receiver
	t.time = time.Now()
	t.amount = amount
	t.hash = t.calculateHash()
	return t
}

func (t *transaction) calculateHash() string {
	var hashString string
	var hash [32]byte

	hashString = t.sender + t.receiver + t.time.Format(time.ANSIC)
	hash = sha256.Sum256([]byte(hashString))
	return string(hash[:])
}

func (t *transaction) String() string {
	return fmt.Sprintf("\tsender: %s\n\treceiver: %s\n\ttime: %s\n\thash: %x\n\tamount: %f\n\n",
		t.sender, t.receiver, t.time.Format(time.ANSIC), t.hash, t.amount)
}
