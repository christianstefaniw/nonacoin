package transactions

type TransactionPool struct {
	pendingTransactions []*Transaction
}

func (t *TransactionPool) AppendTransactions(trans ...*Transaction) bool {
	for _, trans := range trans {
		if trans.GetFromAddress() != nil {
			return false
		}

		if !trans.IsValid() {
			return false
		}
	}

	t.pendingTransactions = append(t.pendingTransactions, trans...)

	return true
}

func (t *TransactionPool) ValidTransactions() bool {
	for _, trans := range t.pendingTransactions {
		if !trans.IsValid() {
			return false
		}
	}
	return true
}
