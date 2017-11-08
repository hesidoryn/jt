package storage

func Delete(key string) error {
	_, ok := storage[key]
	if !ok {
		return ErrorNotFound
	}

	resetTTL(key)
	delete(storage, key)
	return nil
}

func Rename(key, newKey string) error {
	i, ok := storage[key]
	if !ok {
		return ErrorNotFound
	}

	resetTTL(key)
	delete(storage, key)
	storage[newKey] = i
	ttlMap[newKey] = newTicker(newKey)
	return nil
}

func GetTTL(key string) int {
	i, ok := storage[key]
	if !ok {
		return -2
	}

	return i.GetTTL()
}

func GetType(key string) string {
	i, ok := storage[key]
	if !ok {
		return TypeNone
	}

	return i.GetType()
}
