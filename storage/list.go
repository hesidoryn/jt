package storage

import (
	"fmt"
	"strings"
)

// ListItem is struct that contains list.
// It implements Item interface.
type ListItem struct {
	Data []string
	TTL  int
	Type string
}

// GetType returns "+list" for ListItem
func (i *ListItem) GetType() string {
	return i.Type
}

// SetTTL sets time to live value for list item
func (i *ListItem) SetTTL(ttl int) {
	i.TTL = ttl
}

// GetTTL returns time to live value for list item
func (i *ListItem) GetTTL() int {
	return i.TTL
}

func (s *JTStorage) LPush(key, val string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		li := &ListItem{
			Data: []string{val},
			Type: typeList,
			TTL:  -1,
		}
		s.data[key] = li
		return ":1", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return ":0", ErrorWrongType
	}

	li.Data = append(li.Data, val)
	li.Data[0], li.Data[len(li.Data)-1] = li.Data[len(li.Data)-1], li.Data[0]

	res := fmt.Sprintf(":%d", len(li.Data))
	return res, nil
}

func (s *JTStorage) RPush(key, val string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		li := &ListItem{
			Data: []string{val},
			Type: typeList,
			TTL:  -1,
		}
		s.data[key] = li
		return ":1", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return ":0", ErrorWrongType
	}

	li.Data = append(li.Data, val)
	res := fmt.Sprintf(":%d", len(li.Data))
	return res, nil
}

func (s *JTStorage) LPop(key string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return "$-1", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return "$-1", ErrorWrongType
	}

	if len(li.Data) == 0 {
		return "$-1", nil
	}

	var pop string
	pop, li.Data = li.Data[0], li.Data[1:]
	res := fmt.Sprintf("$%d\r\n%s", len(pop), pop)
	return res, nil
}

func (s *JTStorage) RPop(key string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return "$-1", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return "", ErrorWrongType
	}

	if len(li.Data) == 0 {
		return "$-1", nil
	}

	pop := li.Data[len(li.Data)-1]
	li.Data = li.Data[:len(li.Data)-1]
	res := fmt.Sprintf("$%d\r\n%s", len(pop), pop)
	return res, nil
}

func (s *JTStorage) LRem(key string, count int, val string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return "", ErrorWrongType
	}

	newData := make([]string, 0)
	if count == 0 {
		for i := range li.Data {
			if li.Data[i] != val {
				newData = append(newData, li.Data[i])
			}
		}
	}
	if count > 0 {
		for i := range li.Data {
			if li.Data[i] != val {
				newData = append(newData, li.Data[i])
				continue
			}

			if li.Data[i] == val && count > 0 {
				count--
				continue
			}

			newData = append(newData, li.Data[i])
		}
	}
	if count < 0 {
		for i := len(li.Data) - 1; i >= 0; i-- {
			if li.Data[i] != val {
				newData = append([]string{li.Data[i]}, newData...)
				continue
			}

			if li.Data[i] == val && count < 0 {
				count++
				continue
			}

			newData = append([]string{li.Data[i]}, newData...)
		}
	}
	rcount := len(li.Data) - len(newData)
	li.Data = newData

	res := fmt.Sprintf(":%d", rcount)
	return res, nil
}

func (s *JTStorage) LIndex(key string, index int) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return "", ErrorWrongType
	}

	ldata := len(li.Data)
	if index < 0 {
		index += ldata
	}

	if index > ldata || index < 0 {
		return "$-1", nil
	}

	data := li.Data[index]
	res := fmt.Sprintf(":%d\r\n%s", len(data), data)
	return res, nil
}

func (s *JTStorage) LRange(key string, start, end int) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return "*0", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return "", ErrorWrongType
	}

	ldata := len(li.Data)
	if start > ldata {
		return "*0", nil
	}

	if start < 0 {
		start += ldata
	}
	if start < 0 {
		start = 0
	}

	if end > ldata {
		end = ldata - 1
	}
	if end < 0 {
		end += ldata
	}

	if start > end {
		return "*0", nil
	}

	result := []string{}
	for i := start; i <= end; i++ {
		result = append(result, fmt.Sprintf("$%d", len(li.Data[i])), li.Data[i])
	}

	res := fmt.Sprintf("*%d\r\n%s", len(result)/2, strings.Join(result, "\r\n"))
	return res, nil
}

func (s *JTStorage) LLen(key string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0", nil
	}

	li, ok := i.(*ListItem)
	if !ok {
		return "", ErrorWrongType
	}

	res := fmt.Sprintf(":%d", len(li.Data))
	return res, nil
}
