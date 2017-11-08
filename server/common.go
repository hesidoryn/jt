package server

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

func initCommonHandlers() {
	handlers["PING"] = handlerPing
	handlers["DEL"] = handlerDelete
	handlers["RENAME"] = handlerRename
	handlers["TTL"] = handlerTTL
	handlers["PERSIST"] = handlerPersist
	handlers["EXPIRE"] = handlerExpire
	handlers["TYPE"] = handlerType
	handlers["KEYS"] = handlerKeys
	handlers["EXISTS"] = handlerExists
}

func handlerPing(args [][]byte, w *bufio.Writer) {
	if len(args) != 1 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	sendResult(resultPONG, w)
}

func handlerDelete(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	err := storage.Delete(key)
	if err != nil {
		sendResult(":0", w)
		return
	}

	sendResult(":1", w)
}

func handlerRename(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	newKey := string(args[2])
	err := storage.Rename(key, newKey)
	if err != nil {
		sendResult(errorNoSuchKey, w)
		return
	}

	sendResult(resultOK, w)
}

func handlerTTL(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	ttl := storage.GetTTL(key)

	res := fmt.Sprintf(":%d", ttl)
	sendResult(res, w)
}

func handlerPersist(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	res := storage.Persist(key)
	sendResult(res, w)
}

func handlerExpire(args [][]byte, w *bufio.Writer) {
	if len(args) != 3 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	ttl, err := strconv.Atoi(string(args[2]))
	if err != nil {
		sendResult(errorIsNotInteger, w)
		return
	}

	err = storage.SetExpiration(key, ttl)
	if err != nil {
		sendResult(":0", w)
		return
	}

	sendResult(":1", w)
}

func handlerType(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	key := string(args[1])
	res := storage.GetType(key)
	sendResult(res, w)
}

func handlerKeys(args [][]byte, w *bufio.Writer) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	search := string(args[1])
	res := storage.Keys(search)
	sendResult(res, w)
}

func handlerExists(args [][]byte, w *bufio.Writer) {
	if len(args) < 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), w)
		return
	}

	keys := []string{}
	for i := 1; i < len(args); i++ {
		keys = append(keys, string(args[i]))
	}
	res := storage.Exists(keys)
	sendResult(res, w)
}
