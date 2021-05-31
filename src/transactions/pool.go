package transactions

import "bytes"

type TransactionPool struct {
	pendingTransactions []*Transaction
	threshold           uint8
}

func NewPool(threshold uint8) *TransactionPool {
	return &TransactionPool{
		pendingTransactions: make([]*Transaction, 0),
		threshold:           threshold,
	}
}

func (t *TransactionPool) CheckThreshold() uint8 {
	return uint8(len(t.pendingTransactions)) - t.threshold
}

func (t *TransactionPool) AppendTransactions(transactions ...*Transaction) bool {
	for _, trans := range transactions {
		if trans.GetFromAddress() != nil {
			return false
		}

		if !trans.IsValid() {
			return false
		}
	}

	t.pendingTransactions = append(t.pendingTransactions, transactions...)

	return true
}

func (t *TransactionPool) TransactionExists(trans *Transaction) bool {
	for _, pendingTrans := range t.pendingTransactions {
		if bytes.Equal(pendingTrans.hash, trans.hash) {
			return true
		}
	}
	return false
}

func (t *TransactionPool) ValidTransactions() bool {
	for _, trans := range t.pendingTransactions {
		if !trans.IsValid() {
			return false
		}
	}
	return true
}
