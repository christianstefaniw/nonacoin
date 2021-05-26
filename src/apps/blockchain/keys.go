package blockchain

import (
	"crypto/ed25519"
	"crypto/rand"
)

type PublicKey ed25519.PublicKey
type PrivateKey ed25519.PrivateKey

func (pubKey PublicKey) toString() string {
	return string(pubKey)
}

func (privKey PrivateKey) toString() string {
	return string(privKey)
}

func genKeys() (PublicKey, PrivateKey, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	return PublicKey(pubKey), PrivateKey(privKey), err
}
