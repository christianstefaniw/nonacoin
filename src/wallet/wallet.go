package wallet

import (
	"fmt"
	"nonacoin/src/crypto"
)

type Wallet struct {
	PubKey  crypto.PublicKey  `json:"pub key"`
	PrivKey crypto.PrivateKey `json:"priv key"`
}

func NewWallet() *Wallet {
	w := new(Wallet)
	w.PubKey, w.PrivKey, _ = crypto.GenKeys()
	return w
}

func (w *Wallet) GetPubKey() crypto.PublicKey {
	return w.PubKey
}

func (w *Wallet) GetPrivKey() crypto.PrivateKey {
	return w.PrivKey
}

func (w *Wallet) String() string {
	out := fmt.Sprintf("public key: %s\nprivate key: %s", w.PubKey, w.PrivKey)
	return out
}
