package storage

import (
	"time"
)

var ttlCheckers = map[string]chan bool{}

func newTicker(key string) chan bool {
	done := make(chan bool, 1)
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				ttl := storage[key].GetTTL()
				if ttl > 0 {
					storage[key].SetTTL(ttl - 1)
					continue
				}

				if ttl == 0 {
					ticker.Stop()
					delete(ttlCheckers, key)
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

func startTTLChecker(key string) {
	ttlCheckers[key] = newTicker(key)
}

func stopTTLChecker(key string) {
	done, ok := ttlCheckers[key]
	if !ok {
		return
	}

	done <- true
	delete(ttlCheckers, key)
}
