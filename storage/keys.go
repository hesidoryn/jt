package storage

import (
	"path"
)

func (s *JTStorage) Delete(key string) int {
	s.l.Lock()
	defer s.l.Unlock()

	_, ok := s.data[key]
	if !ok {
		return 0
	}

	s.stopTTLChecker(key)
	delete(s.data, key)
	return 1
}

func (s *JTStorage) Rename(key, newKey string) error {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ErrorIsNotExist
	}

	s.stopTTLChecker(key)

	delete(s.data, key)
	s.data[newKey] = i

	s.startTTLChecker(newKey)
	return nil
}

func (s *JTStorage) Persist(key string) int {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return 0
	}

	if i.GetTTL() != -1 {
		s.stopTTLChecker(key)
		i.SetTTL(-1)
		return 1
	}

	return 0
}

func (s *JTStorage) Expire(key string, ttl int) int {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return 0
	}

	i.SetTTL(ttl)
	s.startTTLChecker(key)
	return 1
}

func (s *JTStorage) GetTTL(key string) int {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return -2
	}

	return i.GetTTL()
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

func (s *JTStorage) Keys(pattern string) []string {
	s.l.Lock()
	defer s.l.Unlock()

	result := []string{}
	for key := range s.data {
		ok, err := path.Match(pattern, key)
		if err != nil {
			continue
		}

		if ok {
			result = append(result, key)
		}
	}

	return result
}

func (s *JTStorage) Exists(keys []string) int {
	s.l.Lock()
	defer s.l.Unlock()

	count := 0
	for i := range keys {
		_, ok := s.data[keys[i]]
		if ok {
			count++
		}
	}

	return count
}
