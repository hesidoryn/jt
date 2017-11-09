package client

import (
	"bufio"
	"fmt"
)

func (c *JTClient) Set(key, val string) error {
	command := fmt.Sprintf("SET \"%s\" \"%s\"\n", key, val)
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		ln := scanner.Bytes()
		if string(ln) == "+OK" {

		}
	}

	return err
}
