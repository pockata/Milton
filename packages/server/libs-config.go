package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Server struct {
		Address string `json:"address"`
	} `json:"server"`

	MQTT struct {
		Server   string `json:"server"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mqtt"`
}

func readConfig() Configuration {
	configFile, fileErr := os.Open("./config.json")

	if fileErr != nil {
		panic("Couldn't find `config.json` file. Copy and rename `config.example.json` to `config.json`")
	}

	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		panic("Couldn't load `config.json` due to syntax error")
	}

	return configuration
}
