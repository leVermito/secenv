package ciphers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
)

// EncryptRSA : Encrypt Data with RSA public key
func EncryptRSA(publicKey *rsa.PublicKey, data []byte) []byte {
	hash := sha512.New()
	encryptedData, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, data, nil)
	if err != nil {
		panic(err.Error())
	}

	return encryptedData
}

// DecryptRSA : Decrypt Data with RSA private key
func DecryptRSA(privateKey *rsa.PrivateKey, data []byte) []byte {
	hash := sha512.New()
	data, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, data, nil)
	if err != nil {
		panic(err.Error())
	}

	return data
}

// GeneratePrivateKey : generate rsa private key of given size
func GeneratePrivateKey(size int) *rsa.PrivateKey {
	unprotectedPrivateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		panic(err.Error())
	}

	return unprotectedPrivateKey
}

// EncodePrivateKey : encode unsealed private key
func EncodePrivateKey(privateKey *rsa.PrivateKey) []byte {

	encodedPrivateKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	return pem.EncodeToMemory(encodedPrivateKey)
}

// DecodePrivateKey : decode unsealed private key to rsa.PrivateKey
func DecodePrivateKey(privateKey []byte) *rsa.PrivateKey {

	decodedPrivateKey, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		panic(err.Error())
	}

	return decodedPrivateKey
}

// EncryptEncodedPrivateKey : seal PEM formatted private key with password
func EncryptEncodedPrivateKey(privateKey []byte, keyPassword []byte) []byte {
	block, _ := pem.Decode(privateKey)
	protectedPrivateKey, err := x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, keyPassword, x509.PEMCipherAES256)
	if err != nil {
		panic(err.Error())
	}

	return pem.EncodeToMemory(protectedPrivateKey)
}

// DecryptEncodedPrivateKey : unseal PEM formatted private key with password
func DecryptEncodedPrivateKey(privateKey []byte, keyPassword []byte) []byte {

	protectedPrivateKey, _ := pem.Decode(privateKey)
	unprotectedPrivateKey, err := x509.DecryptPEMBlock(protectedPrivateKey, keyPassword)
	if err != nil {
		panic(err.Error())
	}

	return unprotectedPrivateKey
}
