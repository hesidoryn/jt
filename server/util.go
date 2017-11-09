package server

import (
	"bufio"
	"errors"
)

func sendResult(val string, w *bufio.Writer) {
	w.WriteString(val)
	w.WriteString("\r\n")
	w.Flush()
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
