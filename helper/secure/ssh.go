package secure

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

func NewSSHKeypair() (string, string, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	publicKeyBytes, err := NewSSHPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	return string(privateKeyBytes), string(publicKeyBytes), nil

}

func NewSSHPublicKey(privatekey *rsa.PublicKey) ([]byte, error) {

	publicKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	publicKeyBytes := ssh.MarshalAuthorizedKey(publicKey)

	return publicKeyBytes, nil

}
