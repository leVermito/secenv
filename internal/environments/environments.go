package environments

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func readKeyFromStdin() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter secret environment key: ")
	if text, err := reader.ReadBytes('\n'); err != nil {
		panic(err.Error())
	} else {
		text = bytes.TrimRight(text, "\r\n")
		return text
	}
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
