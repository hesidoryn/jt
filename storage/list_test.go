package storage

import (
	"testing"
)

func TestLPop(t *testing.T) {
	s := Init("")
	key := "lpop"
	// key doesn't exist in storage
	res, err := s.LPop(key)
	if err != ErrorIsNotExist {
		t.Error("Expected ErrorIsNotExist, got ", err)
	}
	if res != "" {
		t.Error("Expected empty string, got ", res)
	}

	// test key isn't ListItem
	s.data[key] = &StringItem{}
	res, err = s.LPop(key)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "" {
		t.Error("Expected empty string, got ", res)
	}

	// key is empty ListItem
	s.data[key] = &ListItem{}
	res, err = s.LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "" {
		t.Error("Expected empty string, got ", res)
	}

	// test key is ListItem
	testData := []string{"item1", "item2"}
	expected := testData[0]
	s.data[key] = &ListItem{Data: testData}
	res, err = s.LPop(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}
