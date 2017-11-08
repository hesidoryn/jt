package storage

import (
	"fmt"
	"strconv"
	"strings"
)

type DictItem struct {
	Data         map[string]string
	TTL          int
	Type         string
	isPersistent int
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

func (i *DictItem) SetPersistence() {
	i.isPersistent = 1
}

func (i *DictItem) GetPersistence() int {
	return i.isPersistent
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
		return "*1\n$-1", nil
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

	result := fmt.Sprintf("*%d\n%s", fcount, strings.Join(res, "\n"))
	return result, nil
}

func DDel(key, field string) (int, error) {
	i, ok := storage[key]
	if !ok {
		return 0, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	delete(di.Data, field)
	return 1, nil
}

func DExists(key, field string) (int, error) {
	i, ok := storage[key]
	if !ok {
		return 0, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	_, ok = di.Data[field]
	if !ok {
		return 0, nil
	}

	return 1, nil
}

func DLen(key string) (int, error) {
	i, ok := storage[key]
	if !ok {
		return 0, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	return len(di.Data), nil
}

func DIncrBy(key, field string, by int) (int, error) {
	i, ok := storage[key]
	if !ok {
		d := map[string]string{field: strconv.Itoa(by)}
		di := &DictItem{
			Data: d,
			Type: TypeDict,
			TTL:  -1,
		}
		storage[key] = di
		return by, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.Itoa(by)
		return by, nil
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, ErrorIsNotInteger
	}

	res := valInt + by
	di.Data[field] = strconv.Itoa(res)
	return res, nil
}

func DIncrByFloat(key, field string, by float64) (float64, error) {
	i, ok := storage[key]
	if !ok {
		d := map[string]string{field: strconv.FormatFloat(by, 'f', -1, 64)}
		di := &DictItem{
			Data: d,
			Type: TypeDict,
			TTL:  -1,
		}
		storage[key] = di
		return by, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.FormatFloat(by, 'f', -1, 64)
		return by, nil
	}

	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, ErrorIsNotFloat
	}

	res := valFloat + by
	di.Data[field] = strconv.FormatFloat(res, 'f', -1, 64)
	return res, nil
}
