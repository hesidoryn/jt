package server

import (
	"strconv"
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

func del(c jtContext) {
	key := string(c.args[1])
	data := jtStorage.Delete(key)
	c.sendResult(prepareIntegerResult(data))
}

func rename(c jtContext) {
	key, newKey := string(c.args[1]), string(c.args[2])
	err := jtStorage.Rename(key, newKey)
	if err != nil {
		c.sendResult(errorNoSuchKey)
		return
	}

	c.sendResult(resultOK)
}

func ttl(c jtContext) {
	key := string(c.args[1])
	data := jtStorage.GetTTL(key)
	c.sendResult(prepareIntegerResult(data))
}

func persist(c jtContext) {
	key := string(c.args[1])
	data := jtStorage.Persist(key)
	c.sendResult(prepareIntegerResult(data))
}

func expire(c jtContext) {
	key := string(c.args[1])
	ttl, err := strconv.Atoi(string(c.args[2]))
	if err != nil {
		c.sendResult(errorIsNotInteger)
		return
	}

	data := jtStorage.Expire(key, ttl)
	c.sendResult(prepareIntegerResult(data))
}

func hType(c jtContext) {
	key := string(c.args[1])
	res := jtStorage.GetType(key)
	c.sendResult(res)
}

func keys(c jtContext) {
	search := string(c.args[1])
	data := jtStorage.Keys(search)
	c.sendResult(prepareListResult(data))
}

func exists(c jtContext) {
	keys := []string{}
	for i := 1; i < len(c.args); i++ {
		keys = append(keys, string(c.args[i]))
	}
	data := jtStorage.Exists(keys)
	c.sendResult(prepareIntegerResult(data))
}
