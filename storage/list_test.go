package storage

import (
	"fmt"
	"testing"
)

func TestLPop(t *testing.T) {
	key := "lpop"
	// key doesn't exist in storage
	res, err := LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key isn't ListItem
	storage[key] = &StringItem{}
	res, err = LPop(key)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// key is empty ListItem
	storage[key] = &ListItem{}
	res, err = LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key is ListItem
	testData := []string{"item1", "item2"}
	expected := fmt.Sprintf("$%d\r\n%s", len(testData[0]), testData[0])
	storage[key] = &ListItem{Data: testData}
	res, err = LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}
