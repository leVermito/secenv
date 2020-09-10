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
		return
	}

	keyPassword := ReadSecretFromStdin("Enter secret environment key:")
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

// ListEnvironments : List stored environments
func ListEnvironments() {
	database := getDatabaseOrCreateNew()
	for key := range database {
		fmt.Println(key)
	}
}

// ShowEnvironment : List variables in environment
func ShowEnvironment(environmentName string, showValues bool) {

	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; !exists {
		fmt.Printf("Environment %s does not exists!\n", environmentName)
		return
	}

	environment := database[environmentName]
	keyPassword := ReadSecretFromStdin("Enter secret environment key:")
	privateKey := ciphers.DecodePrivateKey(ciphers.DecryptEncodedPrivateKey(environment.PrivateKey, keyPassword))
	aesKey := ciphers.DecryptRSA(privateKey, environment.Aes)
	nonce := ciphers.DecryptRSA(privateKey, environment.Nonce)
	data := decodeEnvironmentData(ciphers.DecryptAESGCM(aesKey, nonce, environment.Data))

	if showValues {
		for key := range data {
			fmt.Printf("%s: %s\n", key, data[key].Value)
		}
	} else {
		for key := range data {
			fmt.Printf("%s\n", key)
		}
	}
}

// RemoveEnvironment : Removes the requested environment from database
func RemoveEnvironment(environmentName string) {
	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; !exists {
		fmt.Printf("Environment: %s does not exists!\n", environmentName)
		return
	}

	delete(database, environmentName)
	saveDatabase(database)
}

// AddVariableToEnvironment : Adds variable to specified environment
func AddVariableToEnvironment(environmentName string) {
	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; !exists {
		fmt.Printf("Environment %s does not exists!\n", environmentName)
		return
	}

	environment := database[environmentName]
	keyPassword := ReadSecretFromStdin("Enter secret environment key:")
	privateKey := ciphers.DecodePrivateKey(ciphers.DecryptEncodedPrivateKey(environment.PrivateKey, keyPassword))
	aesKey := ciphers.DecryptRSA(privateKey, environment.Aes)
	nonce := ciphers.DecryptRSA(privateKey, environment.Nonce)
	data := decodeEnvironmentData(ciphers.DecryptAESGCM(aesKey, nonce, environment.Data))

	variable := ReadStringFromStdin("Enter secret variable name: ")

	if _, exists := data[variable]; exists {
		fmt.Printf("Variable: %s already exists!\n", variable)
		return
	}

	value := ReadSecretFromStdin("Enter secret variable value:")
	data[variable] = SecretVariable{"ENV", string(value)}

	sealedEnvironment := sealEnvironment(
		keyPassword,
		UnsealedEnvironment{
			database[environmentName].PrivateKey,
			data,
		},
	)

	database[environmentName] = sealedEnvironment
	saveDatabase(database)
}

// EditVariableFromEnvironment : Edit variable in specified environment
func EditVariableFromEnvironment(environmentName string) {
	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; !exists {
		fmt.Printf("Environment %s does not exists!\n", environmentName)
		return
	}

	environment := database[environmentName]
	keyPassword := ReadSecretFromStdin("Enter secret environment key:")
	privateKey := ciphers.DecodePrivateKey(ciphers.DecryptEncodedPrivateKey(environment.PrivateKey, keyPassword))
	aesKey := ciphers.DecryptRSA(privateKey, environment.Aes)
	nonce := ciphers.DecryptRSA(privateKey, environment.Nonce)
	data := decodeEnvironmentData(ciphers.DecryptAESGCM(aesKey, nonce, environment.Data))

	variable := ReadStringFromStdin("Enter secret variable name: ")

	if _, exists := data[variable]; !exists {
		fmt.Printf("Variable: %s does not exists!\n", variable)
		return
	}

	value := ReadSecretFromStdin("Enter secret variable value:")
	data[variable] = SecretVariable{"ENV", string(value)}

	sealedEnvironment := sealEnvironment(
		keyPassword,
		UnsealedEnvironment{
			database[environmentName].PrivateKey,
			data,
		},
	)

	database[environmentName] = sealedEnvironment
	saveDatabase(database)

}

// RemoveVariableFromEnvironment : Removes variable from specified environment
func RemoveVariableFromEnvironment(environmentName string) {
	database := getDatabaseOrCreateNew()

	if _, exists := database[environmentName]; !exists {
		fmt.Printf("Environment %s does not exists!\n", environmentName)
	}
	environment := database[environmentName]
	keyPassword := ReadSecretFromStdin("Enter secret environment key:")
	privateKey := ciphers.DecodePrivateKey(ciphers.DecryptEncodedPrivateKey(environment.PrivateKey, keyPassword))
	aesKey := ciphers.DecryptRSA(privateKey, environment.Aes)
	nonce := ciphers.DecryptRSA(privateKey, environment.Nonce)
	data := decodeEnvironmentData(ciphers.DecryptAESGCM(aesKey, nonce, environment.Data))

	variable := ReadStringFromStdin("Enter secret variable name: ")

	if _, exists := data[variable]; !exists {
		fmt.Printf("Variable: %s does not exists!\n", variable)
	}

	delete(data, variable)

	sealedEnvironment := sealEnvironment(
		keyPassword,
		UnsealedEnvironment{
			database[environmentName].PrivateKey,
			data,
		},
	)

	database[environmentName] = sealedEnvironment
	saveDatabase(database)
}
