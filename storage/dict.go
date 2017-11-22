package storage

import (
	"fmt"
	"strconv"
	"strings"
)

// DictItem is struct for dict.
// It implements Item interface.
type DictItem struct {
	Data map[string]string
	TTL  int
	Type string
}

// GetType returns "+dict" for DictItem
func (i *DictItem) GetType() string {
	return i.Type
}

// SetTTL sets time to live value for dict item
func (i *DictItem) SetTTL(ttl int) {
	i.TTL = ttl
}

// GetTTL returns time to live value for dict item
func (i *DictItem) GetTTL() int {
	return i.TTL
}

// DSet is used to create dict with some fields
func (s *JTStorage) DSet(key string, vals map[string]string) error {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		di := &DictItem{
			Data: vals,
			Type: typeDict,
			TTL:  -1,
		}
		s.data[key] = di
		return nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return ErrorWrongType
	}

	for k, v := range vals {
		di.Data[k] = v
	}

	return nil
}

// DGet returns expected dict fields
func (s *JTStorage) DGet(key string, fields []string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return "*1\r\n$-1", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	fcount, res := 0, []string{}
	for i := range fields {
		val, ok := di.Data[fields[i]]
		if !ok {
			fcount++
			res = append(res, "$-1")
			continue
		}

		fcount++
		lval := fmt.Sprintf("$%d", len(val))
		res = append(res, lval, val)
	}

	result := fmt.Sprintf("*%d\r\n%s", fcount, strings.Join(res, "\r\n"))
	return result, nil
}

// DDel removes field from dict
func (s *JTStorage) DDel(key, field string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return ":0", ErrorWrongType
	}

	delete(di.Data, field)
	return ":1", nil
}

// DExists checks if field exists in dict
func (s *JTStorage) DExists(key, field string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	_, ok = di.Data[field]
	if !ok {
		return ":0", nil
	}

	return ":1", nil
}

// DLen returns dict's length
func (s *JTStorage) DLen(key string) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return ":0", nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	res := fmt.Sprintf(":%d", len(di.Data))
	return res, nil
}

// DIncrBy increments by "by" value dict's field
// or returns error
func (s *JTStorage) DIncrBy(key, field string, by int) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		d := map[string]string{field: strconv.Itoa(by)}
		di := &DictItem{
			Data: d,
			Type: typeDict,
			TTL:  -1,
		}
		s.data[key] = di

		res := fmt.Sprintf(":%d", by)
		return res, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.Itoa(by)

		res := fmt.Sprintf(":%d", by)
		return res, nil
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return ":0", ErrorIsNotInteger
	}

	newData := valInt + by
	di.Data[field] = strconv.Itoa(newData)

	res := fmt.Sprintf(":%d", newData)
	return res, nil
}

// DIncrByFloat increments by "by" value dict's field
// or returns error
func (s *JTStorage) DIncrByFloat(key, field string, by float64) (string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		val := strconv.FormatFloat(by, 'f', -1, 64)
		d := map[string]string{field: val}
		di := &DictItem{
			Data: d,
			Type: typeDict,
			TTL:  -1,
		}
		s.data[key] = di
		return val, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return "", ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.FormatFloat(by, 'f', -1, 64)
		return di.Data[field], nil
	}

	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return "", ErrorIsNotFloat
	}

	res := valFloat + by
	di.Data[field] = strconv.FormatFloat(res, 'f', -1, 64)
	return di.Data[field], nil
}
