package storage

import "errors"

var storage = map[string]Item{}

type Item interface {
	GetType() string
	SetTTL(int)
	GetTTL() int
	IsPersisted() bool
}

var (
	TypeString = "+string"
	TypeList   = "+list"
	TypeDict   = "+dict"
	TypeNone   = "+none"

	ErrorNotFound     = errors.New("not found")
	ErrorWrongType    = errors.New("wrong type")
	ErrorIsNotInteger = errors.New("value is not integer")
)
