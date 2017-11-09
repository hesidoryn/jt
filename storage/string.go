package storage

import (
	"fmt"
	"strconv"
)

type StringItem struct {
	Data string
	TTL  int
	Type string
}

func (i *StringItem) GetType() string {
	return i.Type
}

func (i *StringItem) SetTTL(ttl int) {
	i.TTL = ttl
}

func (i *StringItem) GetTTL() int {
	return i.TTL
}

func Set(key, val string) {
	resetTTL(key)

	i := &StringItem{
		Data: val,
		Type: TypeString,
		TTL:  -1,
	}
	storage[key] = i
}

func Get(key string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return "$-1", ErrorNotFound
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "$-1", ErrorWrongType
	}

	res := fmt.Sprintf("$%d\r\n%s", len(si.Data), si.Data)
	return res, nil
}

func Append(key, val string) (string, error) {
	i, ok := storage[key]
	if !ok {
		i = &StringItem{
			Data: val,
			Type: TypeString,
			TTL:  -1,
		}
		storage[key] = i
		res := fmt.Sprintf(":%d", len(val))
		return res, nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return ":0", ErrorWrongType
	}

	si.Data += val
	res := fmt.Sprintf(":%d", len(si.Data))
	return res, nil
}

func GetSet(key, val string) (string, error) {
	old, ok := storage[key]
	if !ok {
		i := &StringItem{
			Data: val,
			Type: TypeString,
			TTL:  -1,
		}
		storage[key] = i
		return "$-1", nil
	}

	resetTTL(key)

	if old.GetType() != TypeString {
		return "", ErrorWrongType
	}

	sold := old.(*StringItem)

	new := &StringItem{
		Data: val,
		Type: TypeString,
		TTL:  -1,
	}
	storage[key] = new

	res := fmt.Sprintf("$%d\r\n%s", len(sold.Data), sold.Data)
	return res, nil
}

func Strlen(key string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return "-1", ErrorNotFound
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "", ErrorWrongType
	}

	res := fmt.Sprintf(":%d", len(si.Data))
	return res, nil
}

func IncrBy(key string, by int) (string, error) {
	i, ok := storage[key]
	if !ok {
		i := &StringItem{
			Data: strconv.Itoa(by),
			Type: TypeString,
			TTL:  -1,
		}
		storage[key] = i
		res := fmt.Sprintf(":%d", by)
		return res, nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "", ErrorWrongType
	}

	siInt, err := strconv.Atoi(si.Data)
	if err != nil {
		return "", ErrorIsNotInteger
	}

	si.Data = strconv.Itoa(siInt + by)
	res := fmt.Sprintf(":%s", si.Data)
	return res, nil
}
