package server

import (
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

func lpush(c jtContext) {
	key, val := string(c.args[1]), string(c.args[2])
	data, err := jtStorage.LPush(key, val)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func rpush(c jtContext) {
	key, val := string(c.args[1]), string(c.args[2])
	data, err := jtStorage.RPush(key, val)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func lpop(c jtContext) {
	key := string(c.args[1])
	data, err := jtStorage.LPop(key)
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

func rpop(c jtContext) {
	key := string(c.args[1])
	data, err := jtStorage.RPop(key)
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

func lrem(c jtContext) {
	key := string(c.args[1])
	count, err := strconv.Atoi(string(c.args[2]))
	if err != nil {
		c.sendResult(errorIsNotInteger)
		return
	}
	val := string(c.args[3])
	data, err := jtStorage.LRem(key, count, val)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func lindex(c jtContext) {
	key := string(c.args[1])
	index, err := strconv.Atoi(string(c.args[2]))
	if err != nil {
		c.sendResult(errorIsNotInteger)
		return
	}

	data, err := jtStorage.LIndex(key, index)
	if err == storage.ErrorIsNotExist {
		c.sendResult(prepareIntegerResult(0))
		return
	}
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareStringResult(data))
}

func lrange(c jtContext) {
	key := string(c.args[1])
	start, err := strconv.Atoi(string(c.args[2]))
	if err != nil {
		c.sendResult(errorIsNotInteger)
		return
	}
	end, err := strconv.Atoi(string(c.args[3]))
	if err != nil {
		c.sendResult(errorIsNotInteger)
		return
	}

	data, err := jtStorage.LRange(key, start, end)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareListResult(data))
}

func llen(c jtContext) {
	key := string(c.args[1])
	data, err := jtStorage.LLen(key)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}
