package blockchain

import (
	"testing"
)

func TestCreateBlockchain(t *testing.T) {
	bc := createBlockchain()

	if bc.length() != 1 {
		t.Errorf("blockchain length was incorrect, got: %d, wanted: %d", bc.length(), 1)
	}
}

func TestGenesisBlock(t *testing.T) {
	bc := createBlockchain()

	if bc.getGenesis().GetTransactions() != nil {
		t.Errorf("genisis block transactions was incorrect, got %v, want: %v.", bc.getGenesis().GetTransactions(), nil)
	}
}
