package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

type PublicKey ed25519.PublicKey
type PrivateKey ed25519.PrivateKey

func (pubKey PublicKey) String() string {
	return string(pubKey)
}

func (pubKey PublicKey) Verify(hash, sig []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(pubKey), hash, sig)
}

func (privKey PrivateKey) String() string {
	return string(privKey)
}

func (privKey PrivateKey) Sign(data []byte) []byte {
	return ed25519.Sign(ed25519.PrivateKey(privKey), data)
}

func GenKeys() (PublicKey, PrivateKey, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	return PublicKey(pubKey), PrivateKey(privKey), err
}
