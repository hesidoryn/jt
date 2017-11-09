package server

import (
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

func initKeysHandlers() {
	handlers["DEL"] = handlerDelete
	handlers["RENAME"] = handlerRename
	handlers["TTL"] = handlerTTL
	handlers["PERSIST"] = handlerPersist
	handlers["EXPIRE"] = handlerExpire
	handlers["TYPE"] = handlerType
	handlers["KEYS"] = handlerKeys
	handlers["EXISTS"] = handlerExists
}

func handlerDelete(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	key := string(args[1])
	err := storage.Delete(key)
	if err != nil {
		sendResult(":0", c.w)
		return
	}

	sendResult(":1", c.w)
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
	ttl := storage.GetTTL(key)

	res := fmt.Sprintf(":%d", ttl)
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

	err = storage.SetExpiration(key, ttl)
	if err != nil {
		sendResult(":0", c.w)
		return
	}

	sendResult(":1", c.w)
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