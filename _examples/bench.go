package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hesidoryn/jt/jtclient"
)

const N = 100000

func main() {
	client, err := jtclient.NewClient(jtclient.Options{
		Host:     "localhost:4567",
		Password: "asdf",
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = client.Auth()
	fmt.Println(err)

	// PING_INLINE
	start := time.Now()
	for i := 0; i < N; i++ {
		client.Ping()
	}
	elapsed := time.Since(start)
	rps := strconv.FormatFloat(N/elapsed.Seconds(), 'f', 2, 64)
	fmt.Printf("PING_INLINE: %s requests per second\n", rps)

	// SET
	start = time.Now()
	for i := 0; i < N; i++ {
		client.Set("k", "value")
	}
	elapsed = time.Since(start)
	rps = strconv.FormatFloat(N/elapsed.Seconds(), 'f', 2, 64)
	fmt.Printf("SET: %s requests per second\n", rps)

	// GET
	start = time.Now()
	for i := 0; i < N; i++ {
		client.Get("k")
	}
	elapsed = time.Since(start)
	rps = strconv.FormatFloat(N/elapsed.Seconds(), 'f', 2, 64)
	fmt.Printf("GET: %s requests per second\n", rps)

	// RPUSH
	start = time.Now()
	for i := 0; i < N; i++ {
		client.RPush("list", "value")
	}
	elapsed = time.Since(start)
	rps = strconv.FormatFloat(N/elapsed.Seconds(), 'f', 2, 64)
	fmt.Printf("RPUSH: %s requests per second\n", rps)

	// LPUSH
	start = time.Now()
	for i := 0; i < N; i++ {
		client.LPush("list", "value")
	}
	elapsed = time.Since(start)
	rps = strconv.FormatFloat(N/elapsed.Seconds(), 'f', 2, 64)
	fmt.Printf("LPUSH: %s requests per second\n", rps)

	client.Save()

	// LPOP
	start = time.Now()
	for i := 0; i < N; i++ {
		client.LPop("list")
	}
	elapsed = time.Since(start)
	rps = strconv.FormatFloat(N/elapsed.Seconds(), 'f', 2, 64)
	fmt.Printf("LPOP: %s requests per second\n", rps)

	// RPOP
	start = time.Now()
	for i := 0; i < N; i++ {
		client.RPop("list")
	}
	elapsed = time.Since(start)
	rps = strconv.FormatFloat(N/elapsed.Seconds(), 'f', 2, 64)
	fmt.Printf("RPOP: %s requests per second\n", rps)
}
