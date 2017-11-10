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

	m := map[string]string{
		"field1": "value1",
		"field2": "value2",
	}
	err = client.DSet("dict", m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	resDGet, err := client.DGet("dict", "field1", "field2")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(resDGet)
}
