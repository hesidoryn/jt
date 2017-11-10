package main

import (
	"log"
	"os"

	"github.com/hesidoryn/jt/jtclient"
)

func main() {
	client, err := jtclient.NewClient("client-config.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = client.Set("k", "value")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	resGet, err := client.Get("k")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(resGet)
}
