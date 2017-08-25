package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//LoadConfig attempts to load config based on env and returns a Loaded data object
func LoadConfig(folderPath, baseFileName string) LoadedData {
	env := os.Getenv("ENV")

	if env == "" {
		env = "dev"
	}

	fileName := baseFileName + "." + strings.ToLower(env) + ".json"
	configFileName := folderPath + "/" + fileName

	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Fatal("Unable to load config file from " + configFileName + " because " + err.Error())
	}

	var data map[string]interface{}

	if err = json.Unmarshal(file, &data); err != nil {
		log.Fatal("Unable to load config file from " + configFileName + " because " + err.Error())
	}

	loadedData := loader{fileName, data}
	return &loadedData
}
