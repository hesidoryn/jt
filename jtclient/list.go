package jtclient

import "strconv"

// LPush sends lpush command
func (c *Client) LPush(key string, val string) error {
	_, err := c.sendCommand("LPUSH", key, val)
	return err
}

// RPush sends rpush command
func (c *Client) RPush(key string, val string) error {
	_, err := c.sendCommand("RPUSH", key, val)
	return err
}

// LPop sends lpop command
func (c *Client) LPop(key string) (string, error) {
	res, err := c.sendCommand("LPOP", key)
	return string(res.([]byte)), err
}

// RPop sends rpop command
func (c *Client) RPop(key string) (string, error) {
	res, err := c.sendCommand("RPOP", key)
	return string(res.([]byte)), err
}

// LRange sends lrange command
func (c *Client) LRange(key string, start, end int) ([]string, error) {
	result := []string{}

	args := []string{key, strconv.Itoa(start), strconv.Itoa(end)}
	res, err := c.sendCommand("LRANGE", args...)
	if err != nil {
		return result, err
	}

	ires := res.([][]byte)
	for _, ir := range ires {
		result = append(result, string(ir))
	}

	return result, nil
}
