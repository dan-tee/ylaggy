package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port    int
	Timeout int
}

func readConfig() (config *Config) {
	const configFile = "config.json"
	if _, err := os.Stat(configFile); err != nil {
		panic("Config file not found under ./" + configFile)
	}

	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	config = new(Config)
	if err := decoder.Decode(config); err != nil {
		panic(err)
	}
	return
}
