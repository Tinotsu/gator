// Package config
package config

import (
	"os"
	"encoding/json"
)

type Config struct {
	DBURL string	`json:"db_url"`
	Username string	`json:"current_user_name"`
}

func NewConfig() *Config {
	cfg := Config{}
	return &cfg
}

func Read (cfg *Config) *Config {
	fullPath, err := getConfigFilePath()
	HandleError(err)

	data, err :=  os.ReadFile(fullPath)
	HandleError(err)

	err = json.Unmarshal(data, cfg)
	HandleError(err)
	return cfg
} 

func (cfg *Config) SetUSer (username string) {
	cfg.Username = username

	err := write(*cfg)
	HandleError(err)
}
