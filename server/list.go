package server

import (
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

func initListHandlers() {
	handlers["LPUSH"] = handlerLPush
	handlers["RPUSH"] = handlerRPush
	handlers["LPOP"] = handlerLPop
	handlers["RPOP"] = handlerRPop
	handlers["LREM"] = handlerLRem
	handlers["LINDEX"] = handlerLIndex
	handlers["LRANGE"] = handlerLRange
	handlers["LLEN"] = handlerLLen
}

func handlerLPush(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	l, err := storage.LPush(key, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d", l)
	sendResult(res, c.w)
}

func handlerRPush(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	l, err := storage.RPush(key, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d", l)
	sendResult(res, c.w)
}

func handlerLPop(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res, err := storage.LPop(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerRPop(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res, err := storage.RPop(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerLRem(args [][]byte, c *client) {
	if len(args) != 4 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	count, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, c.w)
		return
	}
	val := string(args[3])
	res, err := storage.LRem(key, count, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerLIndex(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	index, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	res, err := storage.LIndex(key, index)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerLRange(args [][]byte, c *client) {
	if len(args) != 4 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	start, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, c.w)
		return
	}
	end, err := strconv.Atoi(string(args[3]))
	if err != nil {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	res, err := storage.LRange(key, start, end)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerLLen(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res, err := storage.LLen(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}
