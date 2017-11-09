package client

import "fmt"

func (c *JTClient) Rpush(key string, val string) error {
	command := fmt.Sprintf("RPUSH \"%s\" \"%s\"\n", key, val)
	err := c.sendCommand(command)
	if err != nil {
		return err
	}

	_, err = c.readResponse()
	if err != nil {
		return err
	}

	return nil
}
