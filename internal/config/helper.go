package config

import (
	"os"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	HandleError(err)

	fullPath := homePath + "/" + configFileName
	return fullPath, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	HandleError(err)

	content, err := json.Marshal(cfg)
	HandleError(err)
	
	err = os.WriteFile(filePath, content, 0666)
	HandleError(err)

	return nil
}
