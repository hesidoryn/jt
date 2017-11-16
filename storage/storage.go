package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/boltdb/bolt"
)

var locker = &sync.Mutex{}
var storage = map[string]Item{}

// Item is interface that wraps GetType, SetTTL, GetTTL methods
type Item interface {
	GetType() string
	SetTTL(int)
	GetTTL() int
}

const (
	typeString = "+string"
	typeList   = "+list"
	typeDict   = "+dict"
	typeNone   = "+none"

	dumpString = "dumpString"
	dumpList   = "dumpList"
	dumpDict   = "dumpDict"
)

var (
	// ErrorWrongType is sent when command is called for wrong type
	ErrorWrongType = errors.New("wrong type")
	// ErrorIsNotInteger is sent when command is called for non-integer value
	ErrorIsNotInteger = errors.New("value is not integer")
	// ErrorIsNotFloat is sent when command is called for non-float value
	ErrorIsNotFloat = errors.New("value is not float")
)

// Save is used to make backups
func Save() {
	storageCopy := make(map[string]Item)
	for k, v := range storage {
		storageCopy[k] = v
	}

	db, err := bolt.Open(dump, 0600, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		ds, err := tx.CreateBucketIfNotExists([]byte(dumpString))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		dl, err := tx.CreateBucketIfNotExists([]byte(dumpList))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		dd, err := tx.CreateBucketIfNotExists([]byte(dumpDict))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		wg := &sync.WaitGroup{}
		wg.Add(len(storageCopy))
		for k, v := range storageCopy {
			if v.GetType() == typeString {
				go save(k, v, ds, wg)
				continue
			}
			if v.GetType() == typeList {
				go save(k, v, dl, wg)
				continue
			}
			if v.GetType() == typeDict {
				go save(k, v, dd, wg)
				continue
			}
		}
		wg.Wait()

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func save(key string, v Item, b *bolt.Bucket, wg *sync.WaitGroup) {
	defer wg.Done()

	data, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}

	err = b.Put([]byte(key), data)
	if err != nil {
		log.Println(err)
		return
	}
}
