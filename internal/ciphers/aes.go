package ciphers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// GenerateAESKey : Generate 256bits in length AES key
func GenerateAESKey() []byte {
	aesKey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, aesKey); err != nil {
		panic(err.Error())
	}

	return aesKey
}

// GenerateNonce : Generate 96 bits AESGCM nonce
func GenerateNonce() []byte {
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	return nonce
}

// EncryptAESGCM : encrypt data using AES GCM algorithm with provided key and random nonce
func EncryptAESGCM(key []byte, nonce []byte, data []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherData := aesgcm.Seal(nil, nonce, data, nil)

	return cipherData
}

// DecryptAESGCM : decrypt data using AES GCM algorithm with provided key and nonce
func DecryptAESGCM(key []byte, nonce []byte, cipherData []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	data, err := aesgcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		panic(err.Error())
	}

	return data
}
