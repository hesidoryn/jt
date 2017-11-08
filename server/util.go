package server

import (
	"bufio"
)

func sendResult(val string, w *bufio.Writer) {
	w.WriteString(val)
	w.WriteString("\n")
	w.Flush()
}
