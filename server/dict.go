package server

import (
	"strconv"

	"github.com/hesidoryn/jt/storage"
)

const (
	cmdDSet         = "DSET"
	cmdDGet         = "DGET"
	cmdDDel         = "DDEL"
	cmdDExists      = "DEXISTS"
	cmdDLen         = "DLEN"
	cmdDIncrBy      = "DINCRBY"
	cmdDIncrByFloat = "DINCRBYFLOAT"
)

// dset is used for setting dict fields
func dset(c jtContext) {
	key := string(c.args[1])

	vals := make(map[string]string)
	for i := 2; i < len(c.args); i += 2 {
		k, v := string(c.args[i]), string(c.args[i+1])
		vals[k] = v
	}
	err := jtStorage.DSet(key, vals)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(resultOK)
}

func dget(c jtContext) {
	key := string(c.args[1])
	fields := []string{}
	for i := 2; i < len(c.args); i++ {
		fields = append(fields, string(c.args[i]))
	}
	data, err := jtStorage.DGet(key, fields)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareDictResult(data, fields))
}

func ddel(c jtContext) {
	key, field := string(c.args[1]), string(c.args[2])
	data, err := jtStorage.DDel(key, field)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func dexists(c jtContext) {
	key, field := string(c.args[1]), string(c.args[2])
	data, err := jtStorage.DExists(key, field)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func dlen(c jtContext) {
	key := string(c.args[1])
	data, err := jtStorage.DLen(key)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func dincrBy(c jtContext) {
	key, field := string(c.args[1]), string(c.args[2])
	by, err := strconv.Atoi(string(c.args[3]))
	if err != nil {
		c.sendResult(errorIsNotInteger)
		return
	}

	data, err := jtStorage.DIncrBy(key, field, by)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareIntegerResult(data))
}

func dincrByFloat(c jtContext) {
	key, field := string(c.args[1]), string(c.args[2])
	by, err := strconv.ParseFloat(string(c.args[3]), 64)
	if err != nil {
		c.sendResult(errorIsNotFloat)
		return
	}

	data, err := jtStorage.DIncrByFloat(key, field, by)
	if err == storage.ErrorWrongType {
		c.sendResult(errorWrongType)
		return
	}

	c.sendResult(prepareFloatResult(data))
}
