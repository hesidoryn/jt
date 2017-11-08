package storage

import (
	"fmt"
	"math/rand"
	"time"
)

var ttlMap = map[string]chan bool{}

func SetExpiration(key string, ttl int) error {
	i, ok := storage[key]
	if !ok {
		return ErrorNotFound
	}

	i.SetTTL(ttl)
	setNewTTL(key)
	return nil
}

func newTicker(key string) chan bool {
	done := make(chan bool, 1)
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		i := rand.Intn(10)
		for {
			select {
			case <-ticker.C:
				fmt.Println(i)
				ttl := storage[key].GetTTL()
				if ttl > 0 {
					storage[key].SetTTL(ttl - 1)
					continue
				}

				if ttl == 0 {
					ticker.Stop()
					delete(ttlMap, key)
					delete(storage, key)
					return
				}
			case <-done:
				return
			}
		}
	}()

	return done
}

func setNewTTL(key string) {
	ttlMap[key] = newTicker(key)
}

func resetTTL(key string) {
	done, ok := ttlMap[key]
	if !ok {
		return
	}

	done <- true
	delete(ttlMap, key)
}
