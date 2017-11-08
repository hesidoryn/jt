package server

import (
	"bufio"
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

func handlerLPush(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	l, err := storage.LPush(key, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	res := fmt.Sprintf(":%d", l)
	sendResult(res, w)
}

func handlerRPush(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	val := string(args[2])
	l, err := storage.RPush(key, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	res := fmt.Sprintf(":%d", l)
	sendResult(res, w)
}

func handlerLPop(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	res, err := storage.LPop(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	sendResult(res, w)
}

func handlerRPop(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	res, err := storage.RPop(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	sendResult(res, w)
}

func handlerLRem(args [][]byte, w *bufio.Writer) {
	if len(args) != 4 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	count, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, w)
		return
	}
	val := string(args[3])
	res, err := storage.LRem(key, count, val)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	sendResult(res, w)
}

func handlerLIndex(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	index, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, w)
		return
	}

	res, err := storage.LIndex(key, index)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	sendResult(res, w)
}

func handlerLRange(args [][]byte, w *bufio.Writer) {
	if len(args) != 4 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	start, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, w)
		return
	}
	end, err := strconv.Atoi(string(args[3]))
	if err != nil {
		sendResult(errorIsNotInteger, w)
		return
	}

	res, err := storage.LRange(key, start, end)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	sendResult(res, w)
}

func handlerLLen(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	res, err := storage.LLen(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, w)
		return
	}

	sendResult(res, w)
}
