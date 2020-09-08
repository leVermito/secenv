package ciphers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// EncryptRSA : Encrypt Data with RSA public key
func EncryptRSA(publicKey []byte, data []byte) []byte {
	return nil
}

// DecryptRSA : Decrypt Data with RSA private key
func DecryptRSA(keyPassword string, privateKey []byte, data []byte) []byte {
	return nil
}
