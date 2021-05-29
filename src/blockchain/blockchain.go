package blockchain

import (
	"fmt"
	"nonacoin/src/account"
	"nonacoin/src/blocks"
	"nonacoin/src/crypto"
	trans "nonacoin/src/transactions"
	"strings"
)

type blockchain struct {
	account             *account.Account
	difficulty          uint8
	pendingTransactions []*trans.Transaction
	chain               []*blocks.Block
}

func createBlockchain() *blockchain {
	bc := new(blockchain)
	bc.chain = make([]*blocks.Block, 0)
	bc.account = account.GetAccountInstance()
	bc.pendingTransactions = make([]*trans.Transaction, 0)

	bc.chain = append(bc.chain, bc.createGenesisBlock())
	bc.difficulty = 2
	return bc
}

func (bc *blockchain) length() int {
	return len(bc.chain)
}

func (bc *blockchain) getGenesis() *blocks.Block {
	return bc.chain[0]
}

// create the fist block in the blockchain
func (bc *blockchain) createGenesisBlock() *blocks.Block {
	genesis := blocks.CreateBlock(nil, "", 0)
	return genesis
}

// append transactions to the blockchain's pending transactions
func (bc *blockchain) appendTransactions(transactions ...*trans.Transaction) bool {
	for _, trans := range transactions {
		if trans.GetFromAddress() != nil && bc.getBalanceOfWallet(trans.GetFromAddress()) < trans.GetAmount() {
			return false
		}

		if !trans.IsValid() {
			return false
		}
	}
	bc.pendingTransactions = append(bc.pendingTransactions, transactions...)
	return true
}

func (bc *blockchain) appendBlock(b *blocks.Block) {
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
	for i, blck := range bc.chain[1:] {
		currBlock := blck
		prevBlock := bc.chain[i]

		if !currBlock.HasValidTransactions() {
			return false
		}

		if currBlock.GetHash() != currBlock.CalculateHash() {
			return false
		}

		if prevBlock.CalculateHash() != currBlock.GetPrevHash() || prevBlock.GetHash() != currBlock.GetPrevHash() {
			return false
		}
	}

	return true
}

func (bc *blockchain) registerNode() {}

func (bc *blockchain) getBalanceOfWallet(walletAddress crypto.PublicKey) float64 {
	var balance float64

	for _, block := range bc.chain {
		for _, trans := range block.GetTransactions() {
			if trans.GetFromAddress().String() == walletAddress.String() {
				balance -= trans.GetAmount()
			} else if trans.GetToAddress().String() == walletAddress.String() {
				balance += trans.GetAmount()
			}
		}
	}

	return balance
}

func (bc *blockchain) getAllTransactionsForWallet(walletAddress crypto.PublicKey) []*trans.Transaction {
	transactions := make([]*trans.Transaction, 0)

	for _, block := range bc.chain {
		for _, trans := range block.GetTransactions() {
			if trans.GetToAddress().String() == walletAddress.String() || trans.GetFromAddress().String() == walletAddress.String() {
				transactions = append(transactions, trans)
			}
		}
	}

	return transactions
}

func (bc *blockchain) getLatestBlock() *blocks.Block {
	return bc.chain[bc.length()-1]
}
