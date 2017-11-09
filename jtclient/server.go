package jtclient

import (
	"bytes"
	"fmt"
	"io"
)

func (c *Client) Auth() error {
	command := fmt.Sprintf("AUTH \"%s\"\n", c.config.Password)
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, c.conn)
	return err
}
