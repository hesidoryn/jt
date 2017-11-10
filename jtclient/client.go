// Package jtclient provides a client for making commands to jt server.
package jtclient

import (
	"encoding/json"
	"log"
	"net"
	"os"
)

// Client is struct that contains tcp connection to jt server
// and some configuration info.
type Client struct {
	config config
	conn   net.Conn
}

type config struct {
	Host     string `json:"host"`
	Password string `json:"password"`
}

// NewClient creates new jt client
func NewClient(path string) (*Client, error) {
	config := loadConfig(path)
	conn, err := net.Dial("tcp", config.Host)
	if err != nil {
		return &Client{}, err
	}

	c := &Client{
		config: config,
		conn:   conn,
	}
	return c, nil
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
