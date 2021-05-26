package blockchain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strings"
)

func genPrivateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2024)
}

func genPrivatePem(privateKey *rsa.PrivateKey) string {
	var privatePem strings.Builder

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	pem.Encode(&privatePem, privateKeyBlock)

	return privatePem.String()
}

func genPublicPem(publicKey *rsa.PublicKey) (string, error) {
	var publicPem strings.Builder

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	pem.Encode(&publicPem, publicKeyBlock)

	return publicPem.String(), nil
}

// pubKeyFromPem turns a pub formatted key and parses it into a rsa.PublicKey type
func pubKeyFromPem(pubKey string) *rsa.PublicKey {
	pemBlock, _ := pem.Decode([]byte(pubKey))

	parsedKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil
	}

	return parsedKey.(*rsa.PublicKey)
}
