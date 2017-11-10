package jtclient

// Set sends set command
func (c *Client) Set(key, value string) error {
	_, err := c.sendCommand("SET", key, value)
	return err
}

// Get sends get command
func (c *Client) Get(key string) (string, error) {
	res, err := c.sendCommand("GET", key)
	if err != nil {
		return "", err
	}

	return string(res.([]byte)), err
}

// Incr sends incr command
func (c *Client) Incr(key string) (int64, error) {
	res, err := c.sendCommand("INCR", key)
	if err != nil {
		return 0, err
	}

	return res.(int64), err
}
