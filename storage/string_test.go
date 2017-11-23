package storage

import (
	"strconv"
	"testing"

	"github.com/hesidoryn/jt/config"
)

func TestGet(t *testing.T) {
	s := Init(config.Config{})
	key := "get"
	// get key doesn't exist in storage
	res, err := s.Get(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "" {
		t.Error("Expected empty string, got ", res)
	}

	// test key isn't StringItem
	s.data["get"] = &ListItem{}
	res, err = s.Get(key)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "" {
		t.Error("Expected empty string, got ", res)
	}

	// test key is StringItem
	testData := "data"
	expected := testData
	s.data[key] = &StringItem{
		Data: testData,
	}
	res, err = s.Get(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}

func TestAppend(t *testing.T) {
	s := Init(config.Config{})
	key := "append"
	testData := " World!"

	// test key doesn't exist in storage
	expected1 := len(testData)
	res, err := s.Append(key, testData)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected1 {
		t.Error("Expected ", expected1, " got ", res)
	}

	// test key isn't StringItem
	s.data[key] = &ListItem{}
	res, err = s.Append(key, "")
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != 0 {
		t.Error("Expected empty string, got ", res)
	}

	// test key is StringItem
	si := &StringItem{Data: "Hello"}
	s.data[key] = si
	expected2 := len(si.Data) + len(testData)

	res, err = s.Append(key, testData)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected2 {
		t.Error("Expected ", expected2, ", got ", res)
	}
}

func TestGetSet(t *testing.T) {
	s := Init(config.Config{})
	key := "getset"
	testData := " World!"

	// test key doesn't exist in storage
	res, err := s.GetSet(key, testData)
	if err != ErrorIsNotExist {
		t.Error("Expected ErrorIsNotExist, got ", err)
	}
	if res != "" {
		t.Error("Expected empty string, got ", res)
	}

	// test key isn't StringItem
	s.data[key] = &ListItem{}
	res, err = s.GetSet(key, "")
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "" {
		t.Error("Expected empty string, got ", res)
	}

	// test key is StringItem
	oldData := "hello"
	si := &StringItem{Data: oldData}
	s.data[key] = si
	expected := oldData

	res, err = s.GetSet(key, testData)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}

func TestStrlen(t *testing.T) {
	s := Init(config.Config{})
	key := "strlen"
	testData := " World!"

	// test key doesn't exist in storage
	res, err := s.Strlen(key)
	if err != ErrorIsNotExist {
		t.Error("Expected ErrorIsNotExist, got ", err)
	}
	if res != -1 {
		t.Error("Expected -1, got ", res)
	}

	// test key isn't StringItem
	s.data[key] = &ListItem{}
	res, err = s.Strlen(key)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != -1 {
		t.Error("Expected -1, got ", res)
	}

	// test key is StringItem
	si := &StringItem{Data: testData}
	s.data[key] = si
	expected := len(testData)

	res, err = s.Strlen(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}

func TestIncrBy(t *testing.T) {
	s := Init(config.Config{})
	key := "incrby"
	by := 5

	// test key doesn't exist in storage
	expected1 := by
	res, err := s.IncrBy(key, by)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected1 {
		t.Error("Expected ", expected1, " got ", res)
	}

	// test key isn't StringItem
	s.data[key] = &ListItem{}
	res, err = s.IncrBy(key, by)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != 0 {
		t.Error("Expected 0, got ", res)
	}

	// test item isn't integer
	s.data[key] = &StringItem{Data: "sadasd"}
	res, err = s.IncrBy(key, by)
	if err != ErrorIsNotInteger {
		t.Error("Expected ErrorIsNotInteger, got ", err)
	}
	if res != 0 {
		t.Error("Expected 0, got ", res)
	}

	// test item is integer
	siData := 10
	si := &StringItem{Data: strconv.Itoa(siData)}
	s.data[key] = si
	expected2 := siData + by

	res, err = s.IncrBy(key, by)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected2 {
		t.Error("Expected ", expected2, ", got ", res)
	}
}
