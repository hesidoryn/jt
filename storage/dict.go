package storage

import (
	"fmt"
	"strconv"
	"strings"
)

type DictItem struct {
	Data map[string]string
	TTL  int
	Type string
}

func (i *DictItem) GetType() string {
	return i.Type
}

func (i *DictItem) SetTTL(ttl int) {
	i.TTL = ttl
}

func (i *DictItem) GetTTL() int {
	return i.TTL
}

func DSet(key string, vals map[string]string) error {
	i, ok := storage[key]
	if !ok {
		di := &DictItem{
			Data: vals,
			Type: TypeDict,
			TTL:  -1,
		}
		storage[key] = di
		return nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return ErrorWrongType
	}

	for k, v := range vals {
		di.Data[k] = v
	}

	return nil
}

func DGet(key string, fields []string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return "*1\r\n$-1", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	fcount, res := 0, []string{}
	for i := range fields {
		val, ok := di.Data[fields[i]]
		if !ok {
			fcount++
			res = append(res, "$-1")
			continue
		}

		fcount++
		lval := fmt.Sprintf("$%d", len(val))
		res = append(res, lval, val)
	}

	result := fmt.Sprintf("*%d\r\n%s", fcount, strings.Join(res, "\r\n"))
	return result, nil
}

func DDel(key, field string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return ":0", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return ":0", ErrorWrongType
	}

	delete(di.Data, field)
	return ":1", nil
}

func DExists(key, field string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return ":0", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	_, ok = di.Data[field]
	if !ok {
		return ":0", nil
	}

	return ":1", nil
}

func DLen(key string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return ":0", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	res := fmt.Sprintf(":%d", len(di.Data))
	return res, nil
}

func DIncrBy(key, field string, by int) (string, error) {
	i, ok := storage[key]
	if !ok {
		d := map[string]string{field: strconv.Itoa(by)}
		di := &DictItem{
			Data: d,
			Type: TypeDict,
			TTL:  -1,
		}
		storage[key] = di

		res := fmt.Sprintf(":%d", by)
		return res, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.Itoa(by)

		res := fmt.Sprintf(":%d", by)
		return res, nil
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return ":0", ErrorIsNotInteger
	}

	newData := valInt + by
	di.Data[field] = strconv.Itoa(newData)

	res := fmt.Sprintf(":%d", newData)
	return res, nil
}

func DIncrByFloat(key, field string, by float64) (string, error) {
	i, ok := storage[key]
	if !ok {
		val := strconv.FormatFloat(by, 'f', -1, 64)
		d := map[string]string{field: val}
		di := &DictItem{
			Data: d,
			Type: TypeDict,
			TTL:  -1,
		}
		storage[key] = di
		return val, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.FormatFloat(by, 'f', -1, 64)
		return di.Data[field], nil
	}

	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return "", ErrorIsNotFloat
	}

	res := valFloat + by
	di.Data[field] = strconv.FormatFloat(res, 'f', -1, 64)
	return di.Data[field], nil
}
