package environments

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"github.com/Vermibus/secenv/internal/ciphers"
	"io"
)

// secretVariable : containes information about variable itself, its value and category
type secretVariable struct {
	Category string
	Value    string
}

// SealedSecretEnvironment : contains nonce and cipherText of secret environment
type SealedSecretEnvironment struct {
	Nonce []byte
	Data  []byte
}

// UnsealedSecretEnvironment : containes unencrypted data of secret environment
type UnsealedSecretEnvironment struct {
	Data map[string]secretVariable
}

func sealSecretEnviroment(key []byte, secretEnvironment UnsealedSecretEnvironment) SealedSecretEnvironment {

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	var encodedSecretVariable bytes.Buffer
	encoder := gob.NewEncoder(&encodedSecretVariable)
	if err := encoder.Encode(secretEnvironment.Data); err != nil {
		panic(err.Error())
	}

	nonce, encryptedData := ciphers.EncryptAESGCM(key, encodedSecretVariable.Bytes())

	return SealedSecretEnvironment{nonce, encryptedData}
}

func unsealSecretEvironment(key []byte, secretEnvironment SealedSecretEnvironment) UnsealedSecretEnvironment {

	decryptedData := ciphers.DecryptAESGCM(secretEnvironment.Nonce, key, secretEnvironment.Data)

	var encodedSecretVariable = bytes.NewBuffer(decryptedData)
	var decodedSecretVariable map[string]secretVariable

	decoder := gob.NewDecoder(encodedSecretVariable)
	if err := decoder.Decode(&decodedSecretVariable); err != nil {
		panic(err.Error())
	}

	return UnsealedSecretEnvironment{decodedSecretVariable}
}
