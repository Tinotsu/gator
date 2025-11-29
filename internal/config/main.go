// Package config
package config

import (
	"os"
	"log"
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
	if err != nil {log.Fatal(err)}

	data, err :=  os.ReadFile(fullPath)
	if err != nil {log.Fatal(err)}

	err = json.Unmarshal(data, cfg)
	if err != nil {log.Fatal(err)}
	return cfg
} 

func (cfg *Config) SetUSer (username string) {
	cfg.Username = username

	err := write(*cfg)
	if err != nil {
		log.Fatal(err)
	}
}
