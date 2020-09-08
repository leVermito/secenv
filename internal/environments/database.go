package environments

import (
	"encoding/gob"
	"os"
	"os/user"
)

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

func getDatabaseOrCreateNew() map[string]sealedEnvironment {
	mainDirectory := getMainDirectoryOrCreateNew()
	databaseFile := mainDirectory + "/secenv.db"

	if _, err := os.Stat(databaseFile); !os.IsNotExist(err) {
		decodeFile, err := os.Open(databaseFile)
		if err != nil {
			panic(err)
		}
		defer decodeFile.Close()
		decoder := gob.NewDecoder(decodeFile)
		database := make(map[string]sealedEnvironment)
		decoder.Decode(&database)
		return database
	}

	return make(map[string]sealedEnvironment, 0)
}

func saveDatabase(database map[string]sealedEnvironment) {
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
