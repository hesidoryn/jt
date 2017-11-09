package client

func (c *jtclient) Set(key, value string) error {
	_, err := c.sendCommand("SET", key, value)
	return err
}

func (c *jtclient) Get(key string) (string, error) {
	res, err := c.sendCommand("GET", key)
	if err != nil {
		return "", err
	}

	return string(res.([]byte)), err
}
