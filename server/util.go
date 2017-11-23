package server

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type jtContext struct {
	command string
	args    [][]byte
	client  *client
}

func (c *jtContext) sendResult(result string) {
	c.client.w.WriteString(result)
	c.client.w.WriteString("\r\n")
	c.client.w.Flush()
}

type jtHandlerFunc func(c jtContext)

func (s *JTServer) Handle(command string, handler jtHandlerFunc, ms ...jtMiddlewareFunc) {
	result := handler
	for _, m := range ms {
		result = m(result)
	}
	s.routes[command] = result
}

func prepareStringResult(data string) string {
	if len(data) == 0 {
		return resultDefaultString
	}
	return fmt.Sprintf("$%d\r\n%s", len(data), data)
}

func prepareIntegerResult(data int) string {
	return fmt.Sprintf(":%d", data)
}

func prepareFloatResult(data float64) string {
	return strconv.FormatFloat(data, 'f', -1, 64)
}

func prepareListResult(data []string) string {
	if len(data) == 0 {
		return resultDefaultList
	}

	result := []string{}
	for i := 0; i < len(data); i++ {
		result = append(result, fmt.Sprintf("$%d", len(data[i])), data[i])
	}
	return fmt.Sprintf("*%d\r\n%s", len(result)/2, strings.Join(result, "\r\n"))
}

func prepareDictResult(data map[string]string, expected []string) string {
	if len(data) == 0 {
		return resultDefaultDict
	}

	result := []string{}
	for _, f := range expected {
		val, ok := data[f]
		if !ok {
			result = append(result, resultDefaultString)
			continue
		}

		lval := fmt.Sprintf("$%d", len(val))
		result = append(result, lval, val)
	}

	return fmt.Sprintf("*%d\r\n%s", len(expected), strings.Join(result, "\r\n"))
}

func parseArgs(command []byte) ([][]byte, error) {
	var args [][]byte

	var (
		stateStart  = 1
		stateQuotes = 2
		stateArg    = 3
	)

	state := stateStart
	current := make([]byte, 0)

	quote := '"'

	for i := 0; i < len(command); i++ {
		c := command[i]

		if state == stateQuotes {
			if rune(c) != quote {
				current = append(current, c)
				continue
			}

			args = append(args, current)
			current = make([]byte, 0)
			state = stateStart
			continue
		}

		if c == '"' || c == '\'' {
			state = stateQuotes
			quote = rune(c)
			continue
		}

		if state == stateArg {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = make([]byte, 0)
				state = stateStart
				continue
			}

			current = append(current, c)
			continue
		}

		if c != ' ' && c != '\t' {
			state = stateArg
			current = append(current, c)
		}
	}

	if state == stateQuotes {
		return [][]byte{}, errors.New(errorProtocolError)
	}

	if len(current) != 0 {
		args = append(args, current)
	}

	return args, nil
}
