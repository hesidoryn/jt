package client

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func (c *JTClient) sendCommand(command string) error {
	_, err := c.conn.Write([]byte(command))
	return err
}

func (c *JTClient) readResponse() (interface{}, error) {
	reader := bufio.NewReader(c.conn)

	var line string
	var err error

	//read until the first non-whitespace line
	for {
		line, err = reader.ReadString('\n')
		if len(line) == 0 || err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			break
		}
	}

	if line[0] == '+' {
		return strings.TrimSpace(line[1:]), nil
	}

	if line[0] == '-' {
		errmesg := strings.TrimSpace(line[1:])
		return nil, errors.New(errmesg)
	}

	if line[0] == ':' {
		n, err := strconv.ParseInt(strings.TrimSpace(line[1:]), 10, 64)
		if err != nil {
			return nil, errors.New("Int reply is not a number")
		}
		return n, nil
	}

	if line[0] == '*' {
		size, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			return nil, errors.New("MultiBulk reply expected a number")
		}
		if size <= 0 {
			return make([][]byte, 0), nil
		}
		res := make([][]byte, size)
		for i := 0; i < size; i++ {
			res[i], err = readBulk(reader, "")
			if err == errors.New("null") {
				continue
			}
			if err != nil {
				return nil, err
			}
			// dont read end of line as might not have been bulk
		}
		return res, nil
	}
	return readBulk(reader, line)
}

func readBulk(reader *bufio.Reader, head string) ([]byte, error) {
	var err error
	var data []byte

	if head == "" {
		head, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
	}
	switch head[0] {
	case ':':
		data = []byte(strings.TrimSpace(head[1:]))

	case '$':
		size, err := strconv.Atoi(strings.TrimSpace(head[1:]))
		if err != nil {
			return nil, err
		}
		if size == -1 {
			return nil, errors.New("null")
		}
		lr := io.LimitReader(reader, int64(size))
		data, err = ioutil.ReadAll(lr)
		if err == nil {
			// read end of line
			_, err = reader.ReadString('\n')
		}
	default:
		return nil, errors.New("Expecting Prefix '$' or ':'")
	}

	return data, err
}
