package storage

import (
	"errors"
	"log"
	"time"
)

var storage = map[string]Item{}

type Item interface {
	GetType() string
	SetTTL(int)
	GetTTL() int
}

var (
	TypeString = "+string"
	TypeList   = "+list"
	TypeDict   = "+dict"
	TypeNone   = "+none"

	ErrorNotFound     = errors.New("not found")
	ErrorWrongType    = errors.New("wrong type")
	ErrorIsNotInteger = errors.New("value is not integer")
	ErrorIsNotFloat   = errors.New("value is not float")
)

func GetStorage() {
	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, v := range storage {
					log.Println(v)
				}
			}
		}
	}()
}
