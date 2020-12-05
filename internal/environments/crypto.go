package environments

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Vermibus/secenv/internal/ciphers"
	"github.com/Vermibus/secenv/internal/variables"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

// // Variable : asdf
// type Variable struct {
// 	Category string
// 	Value    string
// }

// SealedEnvironment : privateKey, aes, nonce and data are sealed/encrypted
type SealedEnvironment struct {
	PrivateKey []byte // sealed
	Aes        []byte // sealed
	Nonce      []byte // sealed
	Data       []byte // sealed
}

// UnsealedEnvironment : privateKey is selaed, data is decrypted and decoded
type UnsealedEnvironment struct {
	privateKey []byte                        // sealed
	data       map[string]variables.Variable // unsealed
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

// DecryptDataFromSealedEnvironment : Decrypts and returns data from sealed environment
func DecryptDataFromSealedEnvironment(environment SealedEnvironment, keyPassword []byte) map[string]variables.Variable {

	privateKey := ciphers.DecodePrivateKey(ciphers.DecryptEncodedPrivateKey(environment.PrivateKey, keyPassword))
	aesKey := ciphers.DecryptRSA(privateKey, environment.Aes)
	nonce := ciphers.DecryptRSA(privateKey, environment.Nonce)
	data := decodeEnvironmentData(ciphers.DecryptAESGCM(aesKey, nonce, environment.Data))

	return data
}

func encodeEnvironemntData(environmentData map[string]variables.Variable) []byte {
	var encodedEnvironmentData bytes.Buffer
	encoder := gob.NewEncoder(&encodedEnvironmentData)
	if err := encoder.Encode(environmentData); err != nil {
		panic(err.Error())
	}

	return encodedEnvironmentData.Bytes()
}

func decodeEnvironmentData(environmentData []byte) map[string]variables.Variable {
	var decodedEnvironmentData map[string]variables.Variable
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
