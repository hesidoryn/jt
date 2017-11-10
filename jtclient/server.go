package jtclient

// Auth sends auth command
func (c *Client) Auth() error {
	_, err := c.sendCommand("AUTH", c.config.Password)
	return err
}

// Ping sends ping command
func (c *Client) Ping() error {
	_, err := c.sendCommand("PING")
	return err
}

// Save sends save command
func (c *Client) Save() error {
	_, err := c.sendCommand("SAVE")
	return err
}
