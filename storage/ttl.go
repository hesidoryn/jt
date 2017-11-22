package storage

import (
	"time"
)

var ttlCheckers = map[string]chan bool{}

func (s *JTStorage) newTicker(key string) chan bool {
	done := make(chan bool, 1)
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				ttl := s.data[key].GetTTL()
				if ttl > 0 {
					s.data[key].SetTTL(ttl - 1)
					continue
				}

				if ttl == 0 {
					ticker.Stop()
					delete(ttlCheckers, key)
					delete(s.data, key)
					return
				}
			case <-done:
				return
			}
		}
	}()

	return done
}

func (s *JTStorage) startTTLChecker(key string) {
	ttlCheckers[key] = s.newTicker(key)
}

func (s *JTStorage) stopTTLChecker(key string) {
	done, ok := ttlCheckers[key]
	if !ok {
		return
	}

	done <- true
	delete(ttlCheckers, key)
}
