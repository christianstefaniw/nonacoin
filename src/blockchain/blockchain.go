package blockchain

import (
	"fmt"
	"nonacoin/src/account"
	"nonacoin/src/blocks"
	"strings"
)

type Blockchain struct {
	account    *account.Account
	difficulty uint8
	chain      []*blocks.Block
}

func createBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.chain = make([]*blocks.Block, 0)
	bc.account = account.GetAccountInstance()

	bc.chain = append(bc.chain, bc.createGenesisBlock())
	bc.difficulty = 2
	return bc
}

func (bc *Blockchain) length() int {
	return len(bc.chain)
}

func (bc *Blockchain) getGenesis() *blocks.Block {
	return bc.chain[0]
}

// create the fist block in the blockchain
func (bc *Blockchain) createGenesisBlock() *blocks.Block {
	genesis := blocks.CreateBlock(nil, "", 0)
	return genesis
}

func (bc *Blockchain) appendBlock(b *blocks.Block) {
	bc.chain = append(bc.chain, b)
}

func (bc *Blockchain) IsChainValid() bool {
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

func (bc *Blockchain) getLatestBlock() *blocks.Block {
	return bc.chain[bc.length()-1]
}

func (bc *Blockchain) String() string {
	var output strings.Builder

	for _, b := range bc.chain {
		output.WriteString(fmt.Sprint(b))
	}

	output.WriteString(fmt.Sprintf("length: %d", bc.length()))

	return output.String()
}
