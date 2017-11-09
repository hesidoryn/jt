package server

import (
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

const (
	cmdSet    = "SET"
	cmdGet    = "GET"
	cmdAppend = "APPEND"
	cmdGetSet = "GETSET"
	cmdStrlen = "STRLEN"
	cmdIncr   = "INCR"
	cmdIncrBy = "INCRBY"
)

func initStringHandlers() {
	handlers[cmdSet] = handlerSet
	handlers[cmdGet] = handlerGet
	handlers[cmdAppend] = handlerAppend
	handlers[cmdGetSet] = handlerGetSet
	handlers[cmdStrlen] = handlerStrlen
	handlers[cmdIncr] = handlerIncr
	handlers[cmdIncrBy] = handlerIncrBy
}

func handlerSet(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	storage.Set(key, val)
	sendResult(resultOK, c.w)
}

func handlerGet(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res, err := storage.Get(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerAppend(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	append := string(args[2])
	res, err := storage.Append(key, append)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerGetSet(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	res, err := storage.GetSet(key, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerStrlen(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res, err := storage.Strlen(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerIncr(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res, err := storage.IncrBy(key, 1)
	if err == storage.ErrorIsNotInteger {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerIncrBy(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	by, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	val, err := storage.IncrBy(key, by)
	if err == storage.ErrorIsNotInteger {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	sendResult(fmt.Sprintf(":%s", val), c.w)
}
