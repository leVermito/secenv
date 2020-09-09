package environments

import (
	"fmt"
	"github.com/Vermibus/secenv/internal/ciphers"
)

// CreateEnvironment : Creates secret environment
func CreateEnvironment(environmentName string) {

	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; exists {
		fmt.Printf("Environment %s already exists!\n", environmentName)
	} else {
		keyPassword := readKeyFromStdin()
		encryptedPrivateKey := ciphers.EncryptEncodedPrivateKey(
			ciphers.EncodePrivateKey(
				ciphers.GeneratePrivateKey(2048),
			),
			keyPassword,
		)

		sealedEnvironment := sealEnvironment(
			keyPassword,
			UnsealedEnvironment{
				encryptedPrivateKey,
				map[string]SecretVariable{
					"SECENV": SecretVariable{"ENV", environmentName},
				},
			},
		)

		database[environmentName] = sealedEnvironment
		saveDatabase(database)
	}
}

// ListEnvironments : List stored environments
func ListEnvironments() {
	database := getDatabaseOrCreateNew()
	for key := range database {
		fmt.Println(key)
	}
}

// ShowEnvironment : List variables in environment
func ShowEnvironment(environmentName string) {

	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; !exists {
		fmt.Printf("Environment %s does not exists!\n", environmentName)
	} else {
		environment := database[environmentName]
		keyPassword := readKeyFromStdin()
		privateKey := ciphers.DecodePrivateKey(ciphers.DecryptEncodedPrivateKey(environment.PrivateKey, keyPassword))
		aesKey := ciphers.DecryptRSA(privateKey, environment.Aes)
		nonce := ciphers.DecryptRSA(privateKey, environment.Nonce)
		data := decodeEnvironmentData(ciphers.DecryptAESGCM(aesKey, nonce, environment.Data))

		for key := range data {
			fmt.Printf("%s: %s\n", data[key].Category, key)
		}
	}
}
