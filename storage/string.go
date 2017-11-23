package storage

import (
	"strconv"
)

// StringItem is struct that contains string item.
// It imlements Item interface.
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

func (s *JTStorage) Set(key, val string) {
	s.l.Lock()
	defer s.l.Unlock()

	s.stopTTLChecker(key)

	i := &StringItem{
		Data: val,
		Type: typeString,
		TTL:  -1,
	}
	s.data[key] = i
}

func (s *JTStorage) Get(key string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return "", nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "", ErrorWrongType
	}

	return si.Data, nil
}

func (s *JTStorage) Append(key, val string) (int, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		i = &StringItem{
			Data: val,
			Type: typeString,
			TTL:  -1,
		}
		s.data[key] = i
		return len(val), nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return 0, ErrorWrongType
	}

	si.Data += val
	return len(si.Data), nil
}

func (s *JTStorage) GetSet(key, val string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	old, ok := s.data[key]
	if !ok {
		i := &StringItem{
			Data: val,
			Type: typeString,
			TTL:  -1,
		}
		s.data[key] = i
		return "", ErrorIsNotExist
	}

	s.stopTTLChecker(key)

	sold, ok := old.(*StringItem)
	if !ok {
		return "", ErrorWrongType
	}

	new := &StringItem{
		Data: val,
		Type: typeString,
		TTL:  -1,
	}
	s.data[key] = new

	return sold.Data, nil
}

func (s *JTStorage) Strlen(key string) (int, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return -1, ErrorIsNotExist
	}

	si, ok := i.(*StringItem)
	if !ok {
		return -1, ErrorWrongType
	}

	return len(si.Data), nil
}

func (s *JTStorage) IncrBy(key string, by int) (int, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		i := &StringItem{
			Data: strconv.Itoa(by),
			Type: typeString,
			TTL:  -1,
		}
		s.data[key] = i
		return by, nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return 0, ErrorWrongType
	}

	siInt, err := strconv.Atoi(si.Data)
	if err != nil {
		return 0, ErrorIsNotInteger
	}

	result := siInt + by
	si.Data = strconv.Itoa(result)
	return result, nil
}
