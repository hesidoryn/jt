package server

import (
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

func initDictHandlers() {
	handlers["DSET"] = handlerDSet
	handlers["DGET"] = handlerDGet
	handlers["DDEL"] = handlerDDel
	handlers["DEXISTS"] = handlerDExists
	handlers["DLEN"] = handlerDLen
	handlers["DINCRBY"] = handlerDIncrBy
	handlers["DINCRBYFLOAT"] = handlerDIncrByFloat
}

func handlerDSet(args [][]byte, c *client) {
	if len(args) < 4 || len(args)%2 != 0 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])

	vals := make(map[string]string)
	for i := 2; i < len(args); i += 2 {
		k, v := string(args[i]), string(args[i+1])
		vals[k] = v
	}
	err := storage.DSet(key, vals)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(resultOK, c.w)
}

func handlerDGet(args [][]byte, c *client) {
	if len(args) < 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	fields := []string{}
	for i := 2; i < len(args); i++ {
		fields = append(fields, string(args[i]))
	}
	res, err := storage.DGet(key, fields)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerDDel(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	field := string(args[2])
	r, err := storage.DDel(key, field)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d", r)
	sendResult(res, c.w)
}

func handlerDExists(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	field := string(args[2])
	r, err := storage.DExists(key, field)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d", r)
	sendResult(res, c.w)
}

func handlerDLen(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	r, err := storage.DLen(key)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d", r)
	sendResult(res, c.w)
}

func handlerDIncrBy(args [][]byte, c *client) {
	if len(args) != 4 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	field := string(args[2])
	by, err := strconv.Atoi(string(args[3]))
	if err != nil {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	r, err := storage.DIncrBy(key, field, by)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := fmt.Sprintf(":%d", r)
	sendResult(res, c.w)
}

func handlerDIncrByFloat(args [][]byte, c *client) {
	if len(args) != 4 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	field := string(args[2])
	by, err := strconv.ParseFloat(string(args[3]), 64)
	if err != nil {
		sendResult(errorIsNotFloat, c.w)
		return
	}

	r, err := storage.DIncrByFloat(key, field, by)
	if err == storage.ErrorWrongType {
		sendResult(errorWrongType, c.w)
		return
	}

	res := strconv.FormatFloat(r, 'f', -1, 64)
	sendResult(res, c.w)
}
