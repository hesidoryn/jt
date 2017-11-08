package main

import (
	"bytes"
	"errors"
	"fmt"
	"net"
)

type JClient struct {
	Conn net.Conn
}

func main() {
	c := NewClient("127.0.0.1:3333")

	fmt.Println(c.Set("key", "vaue"))
	res, err := c.Get("key")
	fmt.Println(res, err)
	// conn.Write([]byte("get asdfsdf\n"))
	// buff = make([]byte, 1024)
	// n, _ = conn.Read(buff)
	// res = bytes.Split(buff[:n], []byte("\n"))
	// fmt.Println(len(res))
}

func NewClient(addr string) JClient {
	conn, _ := net.Dial("tcp", addr)
	return JClient{
		Conn: conn,
	}
}

func (c *JClient) Set(key, val string) string {
	command := fmt.Sprintf("SET %s %s \n", key, val)
	c.Conn.Write([]byte(command))

	buff := make([]byte, 1024)
	n, _ := c.Conn.Read(buff)

	return string(buff[:n])
}

func (c *JClient) Get(key string) (string, error) {
	command := fmt.Sprintf("GET %s \n", key)
	c.Conn.Write([]byte(command))

	buff := make([]byte, 1024)
	n, _ := c.Conn.Read(buff)
	res := bytes.Split(buff[:n], []byte("\n"))

	if bytes.Equal(res[0], []byte("-1")) {
		return "", errors.New("not found")
	}

	return string(res[1]), nil
}
