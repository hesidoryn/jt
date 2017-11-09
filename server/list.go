package server

import (
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

const (
	cmdLPush  = "LPUSH"
	cmdRPush  = "RPUSH"
	cmdLPop   = "LPOP"
	cmdRPop   = "RPOP"
	cmdLRem   = "LREM"
	cmdLIndex = "LINDEX"
	cmdLRange = "LRANGE"
	cmdLLen   = "LLEN"
)

func initListHandlers() {
	handlers[cmdLPush] = handlerLPush
	handlers[cmdRPush] = handlerRPush
	handlers[cmdLPop] = handlerLPop
	handlers[cmdRPop] = handlerRPop
	handlers[cmdLRem] = handlerLRem
	handlers[cmdLIndex] = handlerLIndex
	handlers[cmdLRange] = handlerLRange
	handlers[cmdLLen] = handlerLLen
}

func handlerLPush(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	res, err := storage.LPush(key, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerRPush(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	res, err := storage.RPush(key, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

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
