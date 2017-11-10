package server

import (
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

const (
	cmdDel     = "DEL"
	cmdRename  = "RENAME"
	cmdTTL     = "TTL"
	cmdPersist = "PERSIST"
	cmdExpire  = "EXPIRE"
	cmdType    = "TYPE"
	cmdKeys    = "KEYS"
	cmdExists  = "EXISTS"
)

func initKeysHandlers() {
	handlers[cmdDel] = handlerDel
	handlers[cmdRename] = handlerRename
	handlers[cmdTTL] = handlerTTL
	handlers[cmdPersist] = handlerPersist
	handlers[cmdExpire] = handlerExpire
	handlers[cmdType] = handlerType
	handlers[cmdKeys] = handlerKeys
	handlers[cmdExists] = handlerExists
}

func handlerDel(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res, err := storage.Delete(key)
	if err != nil {
		sendResult(res, c.w)
		return
	}

	sendResult(res, c.w)
}

func handlerRename(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	newKey := string(args[2])
	err := storage.Rename(key, newKey)
	if err != nil {
		sendResult(errorNoSuchKey, c.w)
		return
	}

	sendResult(resultOK, c.w)
}

func handlerTTL(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res := storage.GetTTL(key)

	sendResult(res, c.w)
}

func handlerPersist(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res := storage.Persist(key)

	sendResult(res, c.w)
}

func handlerExpire(args [][]byte, c *client) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	ttl, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, c.w)
		return
	}

	res := storage.Expire(key, ttl)

	sendResult(res, c.w)
}

func handlerType(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	res := storage.GetType(key)

	sendResult(res, c.w)
}

func handlerKeys(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	search := string(args[1])
	res := storage.Keys(search)

	sendResult(res, c.w)
}

func handlerExists(args [][]byte, c *client) {
	if len(args) < 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	keys := []string{}
	for i := 1; i < len(args); i++ {
		keys = append(keys, string(args[i]))
	}
	res := storage.Exists(keys)

	sendResult(res, c.w)
}
