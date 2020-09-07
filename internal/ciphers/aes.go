package ciphers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// EncryptAESGCM : encrypt data using AES GCM algorithm with provided key and random nonce
func EncryptAESGCM(key []byte, data []byte) ([]byte, []byte) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherData := aesgcm.Seal(nil, nonce, data, nil)

	return nonce, cipherData
}

// DecryptAESGCM : decrypt data using AES GCM algorithm with provided key and nonce
func DecryptAESGCM(nonce []byte, key []byte, cipherData []byte) []byte {

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
