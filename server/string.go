package server

import (
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

func set(c jtContext) {
	key, val := string(c.args[1]), string(c.args[2])
	jtStorage.Set(key, val)
	c.sendResult(resultOK)
}

func get(c jtContext) {
	key := string(c.args[1])
	data, err := jtStorage.Get(key)
	if err == storage.ErrorIsNotExist {
		c.sendResult(resultDefaultString)
		return
	}
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareStringResult(data))
}

func hAppend(c jtContext) {
	key, val := string(c.args[1]), string(c.args[2])
	data, err := jtStorage.Append(key, val)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func getset(c jtContext) {
	key, val := string(c.args[1]), string(c.args[2])
	data, err := jtStorage.GetSet(key, val)
	if err == storage.ErrorIsNotExist {
		c.sendResult(resultDefaultString)
		return
	}
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareStringResult(data))
}

func strlen(c jtContext) {
	key := string(c.args[1])
	data, err := jtStorage.Strlen(key)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func incr(c jtContext) {
	key := string(c.args[1])
	data, err := jtStorage.IncrBy(key, 1)
	if err == storage.ErrorIsNotInteger {
		c.sendResult(errorIsNotInteger)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func incrBy(c jtContext) {
	key := string(c.args[1])
	by, err := strconv.Atoi(string(c.args[2]))
	if err != nil {
		c.sendResult(errorIsNotInteger)
		return
	}

	data, err := jtStorage.IncrBy(key, by)
	if err == storage.ErrorIsNotInteger {
		c.sendResult(errorIsNotInteger)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}
