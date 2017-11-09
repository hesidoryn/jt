// Package jt is a redis-like in-memory database
// It contains server and client parts
//
// By Heorhi Sidoryn
package main

import (
	"flag"
	"fmt"

	"github.com/hesidoryn/jt/config"
	"github.com/hesidoryn/jt/server"
)

func main() {
	configPath := flag.String("config", "", "a string")
	flag.Parse()

	config := config.LoadConfig(*configPath)

	fmt.Println(configPath)
	fmt.Println(config)

	server.Init(config)
}
