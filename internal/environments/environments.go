package environments

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"github.com/Vermibus/secenv/internal/ciphers"
	"io"
	"os"
	"os/user"
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

func getMainDirectoryOrCreateNew() string {

	user, err := user.Current()
	if err != nil {
		panic(err.Error())
	}

	secenvDirectory := user.HomeDir + "/.secenv"

	if _, err := os.Stat(secenvDirectory); os.IsNotExist(err) {
		os.Mkdir(secenvDirectory, os.ModePerm)
	}

	if err != nil {
		panic(err.Error())
	}

	return secenvDirectory
}

func getDatabaseOrCreateNew() map[string]SealedSecretEnvironment {
	mainDirectory := getMainDirectoryOrCreateNew()
	databaseFile := mainDirectory + "/secenv.db"

	if _, err := os.Stat(databaseFile); !os.IsNotExist(err) {
		decodeFile, err := os.Open(databaseFile)
		if err != nil {
			panic(err)
		}
		defer decodeFile.Close()
		decoder := gob.NewDecoder(decodeFile)
		database := make(map[string]SealedSecretEnvironment)
		decoder.Decode(&database)
		return database
	}

	return make(map[string]SealedSecretEnvironment, 0)
}

func saveDatabase(database map[string]SealedSecretEnvironment) {
	mainDirectory := getMainDirectoryOrCreateNew()
	encodeFile, err := os.Create(mainDirectory + "/secenv.db")
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	if err := encoder.Encode(database); err != nil {
		panic(err.Error())
	}
	encodeFile.Close()
}

func readKeyFromStdin() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter secret environment key: ")
	if text, err := reader.ReadBytes('\n'); err != nil {
		panic(err.Error())
	} else {
		return text
	}
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

// CreateEnvironment : Creates secret environment
func CreateEnvironment(environmentName string) {

	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; exists {
		fmt.Printf("Environment %s already exists!\n", environmentName)
	} else {
		key := readKeyFromStdin()
		sealedSecretEnvironment := sealSecretEnviroment(
			key,
			UnsealedSecretEnvironment{
				map[string]secretVariable{
					"SECENV": secretVariable{"ENV", environmentName},
				},
			},
		)

		database[environmentName] = sealedSecretEnvironment
		saveDatabase(database)
	}
}
