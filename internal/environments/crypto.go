package environments

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Vermibus/secenv/internal/ciphers"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

// SecretVariable : containes information about variable itself, its category and value
type SecretVariable struct {
	Category string
	Value    string
}

// SealedEnvironment : privateKey, aes, nonce and data are sealed/encrypted
type SealedEnvironment struct {
	PrivateKey []byte // sealed
	Aes        []byte // sealed
	Nonce      []byte // sealed
	Data       []byte // sealed
}

// UnsealedEnvironment : privateKey is selaed, data is decrypted and decoded
type UnsealedEnvironment struct {
	privateKey []byte                    // sealed
	data       map[string]SecretVariable // unsealed
}

// ReadSecretFromStdin : reads secret from stdin with custom prompt without printing it
func ReadSecretFromStdin(prompt string) []byte {
	fmt.Println(prompt)
	secret, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err.Error())
	}

	return secret
}

// ReadStringFromStdin : reads string from stdin with custom prompt
func ReadStringFromStdin(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n\r")

	return text
}

func encodeEnvironemntData(environmentData map[string]SecretVariable) []byte {
	var encodedEnvironmentData bytes.Buffer
	encoder := gob.NewEncoder(&encodedEnvironmentData)
	if err := encoder.Encode(environmentData); err != nil {
		panic(err.Error())
	}

	return encodedEnvironmentData.Bytes()
}

func decodeEnvironmentData(environmentData []byte) map[string]SecretVariable {
	var decodedEnvironmentData map[string]SecretVariable
	decoder := gob.NewDecoder(bytes.NewBuffer(environmentData))
	if err := decoder.Decode(&decodedEnvironmentData); err != nil {
		panic(err.Error())
	}

	return decodedEnvironmentData
}

func sealEnvironment(keyPassword []byte, environment UnsealedEnvironment) SealedEnvironment {
	aesKey := ciphers.GenerateAESKey()
	nonce := ciphers.GenerateNonce()

	privateKey := ciphers.DecodePrivateKey(
		ciphers.DecryptEncodedPrivateKey(environment.privateKey, keyPassword),
	)

	sealedEnvironment := SealedEnvironment{
		environment.privateKey,
		ciphers.EncryptRSA(&privateKey.PublicKey, aesKey),
		ciphers.EncryptRSA(&privateKey.PublicKey, nonce),
		ciphers.EncryptAESGCM(aesKey, nonce, encodeEnvironemntData(environment.data)),
	}

	return sealedEnvironment
}

func unsealEnvironment(keyPassword []byte, environment SealedEnvironment) UnsealedEnvironment {
	privateKey := ciphers.DecodePrivateKey(
		ciphers.DecryptEncodedPrivateKey(environment.PrivateKey, keyPassword),
	)
	aesKey := ciphers.DecryptRSA(privateKey, environment.Aes)
	nonce := ciphers.DecryptRSA(privateKey, environment.Nonce)

	decodedData := decodeEnvironmentData(
		ciphers.DecryptAESGCM(aesKey, nonce, environment.Data),
	)

	unsealedEnvironment := UnsealedEnvironment{
		environment.PrivateKey,
		decodedData,
	}

	return unsealedEnvironment
}
