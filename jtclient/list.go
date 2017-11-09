package jtclient

func (c *Client) Rpush(key string, val string) error {
	_, err := c.sendCommand("RPUSH", key, val)
	return err
}
