package storage

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

func Delete(key string) (string, error) {
	_, ok := storage[key]
	if !ok {
		return ":0", nil
	}

	resetTTL(key)
	delete(storage, key)
	return ":1", nil
}

func Rename(key, newKey string) error {
	i, ok := storage[key]
	if !ok {
		return errors.New("no such key")
	}

	resetTTL(key)
	delete(storage, key)
	storage[newKey] = i
	ttlMap[newKey] = newTicker(newKey)
	return nil
}

func Persist(key string) string {
	i, ok := storage[key]
	if !ok {
		return ":0"
	}

	if i.GetTTL() != -1 {
		resetTTL(key)
		i.SetTTL(-1)
		return ":1"
	}

	return ":0"
}

func SetExpiration(key string, ttl int) string {
	i, ok := storage[key]
	if !ok {
		return ":0"
	}

	i.SetTTL(ttl)
	setNewTTL(key)
	return ":1"
}

func GetTTL(key string) string {
	i, ok := storage[key]
	if !ok {
		return ":-2"
	}

	res := fmt.Sprintf(":%d", i.GetTTL())
	return res
}

func GetType(key string) string {
	i, ok := storage[key]
	if !ok {
		return typeNone
	}

	return i.GetType()
}

func Keys(pattern string) string {
	res := []string{}
	for key := range storage {
		ok, err := path.Match(pattern, key)
		if err != nil {
			continue
		}

		if ok {
			lkey := fmt.Sprintf("$%d", len(key))
			res = append(res, lkey, key)
		}
	}

	if len(res) == 0 {
		return "*0"
	}

	lres := fmt.Sprintf("*%d\r\n", len(res)/2)
	result := lres + strings.Join(res, "\r\n")
	return result
}

func Exists(keys []string) string {
	count := 0
	for i := range keys {
		_, ok := storage[keys[i]]
		if ok {
			count++
		}
	}

	return fmt.Sprintf(":%d", count)
}
