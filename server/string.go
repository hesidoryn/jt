package server

import (
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

func handlerSet(args [][]byte, c *client) {
	if len(args) != 3 {
		fmt.Println(4)
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
	val, err := storage.Get(key)
	if err != nil {
		sendResult("$-1", c.w)
		return
	}

	res := fmt.Sprintf("$%d\n%s\n", len(val), val)
	sendResult(res, c.w)
	return
}

func handlerAppend(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	append := string(args[2])
	l, err := storage.Append(key, append)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d", l)
	sendResult(res, c.w)
}

func handlerGetSet(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	oldVal, err := storage.Get(key)
	if err != nil {
		sendResult("$-1", c.w)
		return
	}

	val := string(args[2])
	storage.Set(key, val)

	res := fmt.Sprintf("$%d\n%s\n", len(val), oldVal)
	sendResult(res, c.w)
}

func handlerStrlen(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val, err := storage.Strlen(key)
	if err == storage.ErrorNotFound {
		sendResult(":0", c.w)
		return
	}
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d\n", len(val))
	sendResult(res, c.w)
}

func handlerIncr(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val, err := storage.IncrBy(key, 1)
	if err == storage.ErrorIsNotInteger {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	sendResult(fmt.Sprintf(":%s", val), c.w)
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
