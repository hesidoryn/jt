// JT is a Redis-like in-memory database.
//
// It contains server and client parts.
//
// By Heorhi Sidoryn
package main

import (
	"github.com/hesidoryn/jt/config"
	"github.com/hesidoryn/jt/server"
)

func main() {
	config := config.Config{}
	config.Parse()
	server.Init(config)
}
