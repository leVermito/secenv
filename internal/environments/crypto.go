package environments

import (
	"bytes"
	"encoding/gob"
	"github.com/Vermibus/secenv/internal/ciphers"
)

// secretVariable : containes information about variable itself, its category and value
type secretVariable struct {
	Category string
	Value    string
}

// sealedEnvironment :
type sealedEnvironment struct {
	privateKey []byte
	publicKey  []byte
	data       []byte
}

type unsealedEnvironment struct {
	privateKey []byte
	publicKey  []byte
	data       map[string]secretVariable
}

func encodeEnvironemntData(environmentData map[string]secretVariable) []byte {

	var encodedEnvironmentData bytes.Buffer
	encoder := gob.NewEncoder(&encodedEnvironmentData)
	if err := encoder.Encode(environmentData); err != nil {
		panic(err.Error())
	}

	return encodedEnvironmentData.Bytes()
}

func decodeEnvironmentData(environmentData []byte) map[string]secretVariable {

	var decodedEnvironmentData map[string]secretVariable

	decoder := gob.NewDecoder(bytes.NewBuffer(environmentData))
	if err := decoder.Decode(&decodedEnvironmentData); err != nil {
		panic(err.Error())
	}

	return decodedEnvironmentData
}

func sealEnvironment(keyPassword string, environment unsealedEnvironment) sealedEnvironment {

	return sealedEnvironment{
		environment.privateKey,
		environment.publicKey,
		ciphers.EncryptRSA(
			environment.publicKey,
			encodeEnvironemntData(environment.data),
		),
	}
}

func unsealEnvironment(keyPassword string, environment sealedEnvironment) unsealedEnvironment {

	return unsealedEnvironment{
		environment.privateKey,
		environment.publicKey,
		decodeEnvironmentData(
			ciphers.DecryptRSA(
				keyPassword,
				environment.privateKey,
				environment.data,
			),
		),
	}
}
