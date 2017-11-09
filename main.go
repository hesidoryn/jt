// Package jt is a redis-like in-memory database
// It contains server and client parts
//
// By Heorhi Sidoryn
package main

import (
	"flag"

	"github.com/hesidoryn/jt/config"
	"github.com/hesidoryn/jt/server"
)

func main() {
	configPath := flag.String("config", "", "a string")
	flag.Parse()

	config := config.LoadConfig(*configPath)

	server.Init(config)

}
