package client

import "fmt"

func (c *jtclient) DSet(key string, dict map[string]string) error {
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

func (c *jtclient) DGet(key string, args ...string) (string, error) {
	args = append([]string{key}, args...)
	res, err := c.sendCommand("DGET", args)
	fmt.Println(res)
	return string(res.([]byte)), err
}
