package storage

import (
	"fmt"
	"path"
	"strings"
)

func Delete(key string) error {
	_, ok := storage[key]
	if !ok {
		return ErrorNotFound
	}

	resetTTL(key)
	delete(storage, key)
	return nil
}

func Rename(key, newKey string) error {
	i, ok := storage[key]
	if !ok {
		return ErrorNotFound
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

func GetTTL(key string) int {
	i, ok := storage[key]
	if !ok {
		return -2
	}

	return i.GetTTL()
}

func GetType(key string) string {
	i, ok := storage[key]
	if !ok {
		return TypeNone
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

	lres := fmt.Sprintf("*%d\n", len(res)/2)
	result := lres + strings.Join(res, "\n")
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
