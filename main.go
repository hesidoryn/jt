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
)

const (
	defaultPort     = "3333"
	defaultDumpFile = "dump.db"
)

func main() {
	config := config.Config{}
	flag.StringVar(&config.Port, "port", defaultPort, "port for bind server")
	flag.StringVar(&config.Password, "password", "", "password for server")
	flag.StringVar(&config.DumpFile, "dump", defaultDumpFile, "file with saved data")
	flag.Parse()

	server.Init(config)
}
