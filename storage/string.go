package storage

import (
	"fmt"
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
		return "$-1", nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "$-1", ErrorWrongType
	}

	res := fmt.Sprintf("$%d\r\n%s", len(si.Data), si.Data)
	return res, nil
}

func (s *JTStorage) Append(key, val string) (string, error) {
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
		return "$-1", nil
	}

	s.stopTTLChecker(key)

	sold, ok := old.(*StringItem)
	if !ok {
		return "$-1", ErrorWrongType
	}

	new := &StringItem{
		Data: val,
		Type: typeString,
		TTL:  -1,
	}
	s.data[key] = new

	res := fmt.Sprintf("$%d\r\n%s", len(sold.Data), sold.Data)
	return res, nil
}

func (s *JTStorage) Strlen(key string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return "-1", nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "-1", ErrorWrongType
	}

	res := fmt.Sprintf(":%d", len(si.Data))
	return res, nil
}

func (s *JTStorage) IncrBy(key string, by int) (string, error) {
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
		res := fmt.Sprintf(":%d", by)
		return res, nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return ":0", ErrorWrongType
	}

	siInt, err := strconv.Atoi(si.Data)
	if err != nil {
		return ":0", ErrorIsNotInteger
	}

	si.Data = strconv.Itoa(siInt + by)
	res := fmt.Sprintf(":%s", si.Data)
	return res, nil
}
