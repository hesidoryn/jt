package storage

import (
	"fmt"
	"testing"

	"github.com/hesidoryn/jt/config"
)

func TestLPop(t *testing.T) {
	s := Init(config.Config{})
	key := "lpop"
	// key doesn't exist in storage
	res, err := s.LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key isn't ListItem
	s.data[key] = &StringItem{}
	res, err = s.LPop(key)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// key is empty ListItem
	s.data[key] = &ListItem{}
	res, err = s.LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key is ListItem
	testData := []string{"item1", "item2"}
	expected := fmt.Sprintf("$%d\r\n%s", len(testData[0]), testData[0])
	s.data[key] = &ListItem{Data: testData}
	res, err = s.LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}
