package client

func (c *jtclient) Del(key string) (bool, error) {
	res, err := c.sendCommand("DEL", key)
	if err != nil {
		return false, err
	}

	return res.(int64) == 1, nil
}

func (c *jtclient) Exists(key string) (bool, error) {
	res, err := c.sendCommand("EXISTS", key)
	if err != nil {
		return false, err
	}

	return res.(int64) == 1, nil
}
