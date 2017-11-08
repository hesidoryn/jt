package server

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

func initStringHandlers() {
	handlers["SET"] = handlerSet
	handlers["GET"] = handlerGet
	handlers["APPEND"] = handlerAppend
	handlers["GETSET"] = handlerGetSet
	handlers["STRLEN"] = handlerStrlen
	handlers["INCR"] = handlerIncr
	handlers["INCRBY"] = handlerIncrBy
}

func handlerSet(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	storage.Set(key, val)
	sendResult(resultOK, w)
}

func handlerGet(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	val, err := storage.Get(key)
	if err != nil {
		sendResult("$-1", w)
		return
	}

	w.WriteString(fmt.Sprintf("$%d", len(val)))
	w.WriteString("\n")
	w.WriteString(val)
	w.WriteString("\n")
	w.Flush()
	return
}

func handlerAppend(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	append := string(args[2])
	l, err := storage.Append(key, append)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	res := fmt.Sprintf(":%d", l)
	sendResult(res, w)
}

func handlerGetSet(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	oldVal, err := storage.Get(key)
	if err != nil {
		sendResult("$-1", w)
		return
	}

	val := string(args[2])
	storage.Set(key, val)

	w.WriteString(fmt.Sprintf("$%d", len(val)))
	w.WriteString("\n")
	w.WriteString(oldVal)
	w.WriteString("\n")
	w.Flush()
	return
}

func handlerStrlen(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	val, err := storage.Strlen(key)
	if err == storage.ErrorNotFound {
		sendResult(":0", w)
		return
	}
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	w.WriteString(fmt.Sprintf(":%d", len(val)))
	w.WriteString("\n")
	w.Flush()
	return
}

func handlerIncr(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	val, err := storage.IncrBy(key, 1)
	if err == storage.ErrorIsNotInteger {
		sendResult(errorIsNotInteger, w)
		return
	}

	sendResult(fmt.Sprintf(":%s", val), w)
}

func handlerIncrBy(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	by, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, w)
		return
	}

	val, err := storage.IncrBy(key, by)
	if err == storage.ErrorIsNotInteger {
		sendResult(errorIsNotInteger, w)
		return
	}

	sendResult(fmt.Sprintf(":%s", val), w)
}
