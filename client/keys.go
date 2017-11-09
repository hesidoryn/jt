package client

import "fmt"

func (c *JTClient) Exists(key string) (bool, error) {
	command := fmt.Sprintf("EXISTS \"%s\"\n", key)
	err := c.sendCommand(command)
	if err != nil {
		return false, err
	}

	res, err := c.readResponse()
	if err != nil {
		return false, err
	}
	return res.(int64) == 1, nil
}
