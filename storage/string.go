package storage

import "strconv"

type StringItem struct {
	Data         string
	TTL          int
	Type         string
	isPersistent int
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

func (i *StringItem) SetPersistence() {
	i.isPersistent = 1
}

func (i *StringItem) GetPersistence() int {
	return i.isPersistent
}

func Set(key, val string) {
	resetTTL(key)

	i := &StringItem{
		Data: val,
		Type: TypeString,
		TTL:  -1,
	}
	storage[key] = i
}

func Get(key string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return "", ErrorNotFound
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "", ErrorWrongType
	}

	return si.Data, nil
}

func Append(key, val string) (int, error) {
	i, ok := storage[key]
	if !ok {
		i = &StringItem{
			Data: val,
			Type: TypeString,
			TTL:  -1,
		}
		storage[key] = i
		return len(val), nil
	}

	si, ok := i.(*StringItem)
	if !ok {
		return 0, ErrorWrongType
	}

	si.Data += val
	return len(si.Data), nil
}

func GetSet(key, val string) (string, error) {
	old, ok := storage[key]
	if !ok {
		return "", ErrorNotFound
	}

	resetTTL(key)

	if old.GetType() != TypeString {
		return "", ErrorWrongType
	}

	sold := old.(*StringItem)

	new := &StringItem{
		Data: val,
		Type: TypeString,
		TTL:  -1,
	}
	storage[key] = new
	return sold.Data, nil
}

func Strlen(key string) (string, error) {
	i, ok := storage[key]
	if !ok {
		return "-1", ErrorNotFound
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "", ErrorWrongType
	}

	strlen := strconv.Itoa(len(si.Data))
	return strlen, nil
}

func IncrBy(key string, by int) (string, error) {
	i, ok := storage[key]
	if !ok {
		return "", ErrorNotFound
	}

	si, ok := i.(*StringItem)
	if !ok {
		return "", ErrorWrongType
	}

	siInt, err := strconv.Atoi(si.Data)
	if err != nil {
		return "", ErrorIsNotInteger
	}

	si.Data = strconv.Itoa(siInt + by)
	return si.Data, nil
}
