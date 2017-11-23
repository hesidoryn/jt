package main

import (
	"log"
	"os"

	"github.com/hesidoryn/jt/jtclient"
)

func main() {
	client, err := jtclient.NewClient(jtclient.Options{
		Host: "localhost:4567",
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = client.RPush("list", "first")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
