// Package storage implements in-memory store.
//
// Overview
//
// Strings, lists and dicts can be stored.
//
// Protocol is redis compatible except dicts.
//
// >30 commands are realised.
//
// Persistence is realised using boltdb and provides by SAVE command.
//
package storage

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/hesidoryn/jt/config"
)

// dump with jt data
var dump = "dump.db"

// Init inits storage module and load dump db if needed
func Init(config config.Config) {
	if config.DB == "" {
		return
	}

	dump = config.DB
	loadDump()
}

func loadDump() {
	if _, err := os.Open(dump); err != nil {
		return
	}

	db, err := bolt.Open(dump, 0600, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		ds := tx.Bucket([]byte(dumpString))
		dl := tx.Bucket([]byte(dumpList))
		dd := tx.Bucket([]byte(dumpDict))

		wg := &sync.WaitGroup{}
		wg.Add(3)
		go loadStrings(ds, wg)
		go loadLists(dl, wg)
		go loadDicts(dd, wg)
		wg.Wait()

		return err
	})
	if err != nil {
		log.Println(err)
	}
}

func loadStrings(b *bolt.Bucket, wg *sync.WaitGroup) {
	defer wg.Done()
	err := b.ForEach(func(k, v []byte) error {
		si := &StringItem{}
		err := json.Unmarshal(v, si)
		if err != nil {
			return err
		}
		si.TTL = -1
		storage[string(k)] = si
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func loadLists(b *bolt.Bucket, wg *sync.WaitGroup) {
	defer wg.Done()
	err := b.ForEach(func(k, v []byte) error {
		li := &ListItem{}
		err := json.Unmarshal(v, li)
		if err != nil {
			return err
		}
		li.TTL = -1
		storage[string(k)] = li
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func loadDicts(b *bolt.Bucket, wg *sync.WaitGroup) {
	defer wg.Done()
	err := b.ForEach(func(k, v []byte) error {
		di := &DictItem{}
		err := json.Unmarshal(v, di)
		if err != nil {
			return err
		}
		di.TTL = -1
		storage[string(k)] = di
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}
