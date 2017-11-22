package storage

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

func (s *JTStorage) Delete(key string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	_, ok := s.data[key]
	if !ok {
		return ":0", nil
	}

	s.stopTTLChecker(key)
	delete(s.data, key)
	return ":1", nil
}

func (s *JTStorage) Rename(key, newKey string) error {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return errors.New("no such key")
	}

	s.stopTTLChecker(key)

	delete(s.data, key)
	s.data[newKey] = i

	s.startTTLChecker(newKey)
	return nil
}

func (s *JTStorage) Persist(key string) string {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0"
	}

	if i.GetTTL() != -1 {
		s.stopTTLChecker(key)
		i.SetTTL(-1)
		return ":1"
	}

	return ":0"
}

func (s *JTStorage) Expire(key string, ttl int) string {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0"
	}

	i.SetTTL(ttl)
	s.startTTLChecker(key)
	return ":1"
}

func (s *JTStorage) GetTTL(key string) string {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":-2"
	}

	res := fmt.Sprintf(":%d", i.GetTTL())
	return res
}

func (s *JTStorage) GetType(key string) string {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return typeNone
	}

	return i.GetType()
}

func (s *JTStorage) Keys(pattern string) string {
	s.l.Lock()
	defer s.l.Unlock()

	res := []string{}
	for key := range s.data {
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

func (s *JTStorage) Exists(keys []string) string {
	s.l.Lock()
	defer s.l.Unlock()

	count := 0
	for i := range keys {
		_, ok := s.data[keys[i]]
		if ok {
			count++
		}
	}

	return fmt.Sprintf(":%d", count)
}
