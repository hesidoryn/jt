package storage

import (
	"strconv"
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
func (s *JTStorage) DGet(key string, fields []string) (map[string]string, error) {
	s.l.Lock()
	defer s.l.Unlock()

	r := make(map[string]string)
	i, ok := s.data[key]
	if !ok {
		return r, ErrorIsNotExist
	}

	di, ok := i.(*DictItem)
	if !ok {
		return r, ErrorWrongType
	}

	for _, f := range fields {
		val, ok := di.Data[f]
		if !ok {
			continue
		}

		r[f] = val
	}

	return r, nil
}

// DDel removes field from dict
func (s *JTStorage) DDel(key, field string) (int, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return 0, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	delete(di.Data, field)
	return 1, nil
}

// DExists checks if field exists in dict
func (s *JTStorage) DExists(key, field string) (int, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return 0, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	_, ok = di.Data[field]
	if !ok {
		return 0, nil
	}

	return 1, nil
}

// DLen returns dict's length
func (s *JTStorage) DLen(key string) (int, error) {
	s.l.Lock()
	defer s.l.Unlock()

	i, ok := s.data[key]
	if !ok {
		return 0, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	return len(di.Data), nil
}

// DIncrBy increments by "by" value dict's field
// or returns error
func (s *JTStorage) DIncrBy(key, field string, by int) (int, error) {
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

		return by, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.Itoa(by)
		return by, nil
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, ErrorIsNotInteger
	}

	newData := valInt + by
	di.Data[field] = strconv.Itoa(newData)

	return newData, nil
}

// DIncrByFloat increments by "by" value dict's field
// or returns error
func (s *JTStorage) DIncrByFloat(key, field string, by float64) (float64, error) {
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
		return by, nil
	}

	di, ok := i.(*DictItem)
	if !ok {
		return 0, ErrorWrongType
	}

	val, ok := di.Data[field]
	if !ok {
		di.Data[field] = strconv.FormatFloat(by, 'f', -1, 64)
		return by, nil
	}

	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, ErrorIsNotFloat
	}

	newData := valFloat + by
	di.Data[field] = strconv.FormatFloat(newData, 'f', -1, 64)
	return newData, nil
}
