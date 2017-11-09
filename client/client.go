package client

import (
	"encoding/json"
	"log"
	"net"
	"os"
)

type JTClient struct {
	config config
	conn   net.Conn
}

type config struct {
	Host     string `json:"host"`
	Password string `json:"password"`
}

func loadConfig(path string) config {
	c := config{}
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dec := json.NewDecoder(file)
	err = dec.Decode(&c)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return c
}

func NewClient(path string) (*JTClient, error) {
	config := loadConfig(path)
	conn, err := net.Dial("tcp", config.Host)
	if err != nil {
		return &JTClient{}, err
	}

	client := &JTClient{
		config: config,
		conn:   conn,
	}
	return client, nil
}
