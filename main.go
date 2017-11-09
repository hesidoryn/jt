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
