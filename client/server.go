package client

import (
	"bytes"
	"fmt"
	"io"
)

func (c *JTClient) Auth() error {
	command := fmt.Sprintf("AUTH \"%s\"\n", c.config.Password)
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, c.conn)
	fmt.Println(buf.String())
	return err
}