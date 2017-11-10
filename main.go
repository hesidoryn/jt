// JT is a Redis-like in-memory database.
//
// It contains server and client parts.
//
// By Heorhi Sidoryn
package main

import (
	"flag"

	"github.com/hesidoryn/jt/config"
	"github.com/hesidoryn/jt/server"
	"github.com/hesidoryn/jt/storage"
)

func main() {
	configPath := flag.String("config", "", "a string")
	flag.Parse()

	config := config.LoadConfig(*configPath)
	storage.Init(config)
	server.Init(config)
}
