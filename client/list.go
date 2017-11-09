package client

func (c *jtclient) Rpush(key string, val string) error {
	_, err := c.sendCommand("RPUSH", key, val)
	if err != nil {
		return err
	}

	return err

}
