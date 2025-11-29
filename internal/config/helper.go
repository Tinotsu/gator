package config

import (
	"os"
	"log"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	fullPath := homePath + "/" + configFileName
	return fullPath, err
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {log.Fatal(err)}

	print(filePath,"\n")
	content, err := json.Marshal(cfg)
	print(content,"\n")

	if err != nil {log.Fatal(err)}

	err = os.WriteFile(filePath, content, 0666)
	return err
}
