package jtclient

// DSet creates dict with specific fields and values
func (c *Client) DSet(key string, dict map[string]string) error {
	args := make([]string, len(dict)*2+1)
	args[0] = key
	i := 1
	for k, v := range dict {
		args[i] = k
		args[i+1] = string(v)
		i += 2
	}

	_, err := c.sendCommand("DSET", args...)
	return err
}

// DGet is used to get dict's fields
func (c *Client) DGet(key string, args ...string) ([]string, error) {
	result := []string{}

	args = append([]string{key}, args...)
	res, err := c.sendCommand("DGET", args...)
	if err != nil {
		return result, err
	}

	ires := res.([][]byte)
	for _, ir := range ires {
		result = append(result, string(ir))
	}

	return result, err
}
