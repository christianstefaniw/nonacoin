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

func (pubKey PublicKey) verify(hash, sig []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(pubKey), hash, sig)
}

func (privKey PrivateKey) toString() string {
	return string(privKey)
}

func (privKey PrivateKey) sign(data []byte) []byte {
	return ed25519.Sign(ed25519.PrivateKey(privKey), data)
}

func genKeys() (PublicKey, PrivateKey, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	return PublicKey(pubKey), PrivateKey(privKey), err
}
