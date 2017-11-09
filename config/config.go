package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func LoadConfig(path string) Config {
	config := Config{}

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dec := json.NewDecoder(file)
	err = dec.Decode(&config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return config
}
