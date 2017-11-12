package storage

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	key := "get"
	// get key doesn't exist in storage
	res, err := Get(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key isn't StringItem
	storage["get"] = &ListItem{}
	res, err = Get(key)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key is StringItem
	testData := "data"
	expected := fmt.Sprintf("$%d\r\n%s", len(testData), testData)
	storage[key] = &StringItem{
		Data: testData,
	}
	res, err = Get(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}

func TestAppend(t *testing.T) {
	key := "append"
	testData := " World!"

	// test key doesn't exist in storage
	expected1 := fmt.Sprintf(":%d", len(testData))
	res, err := Append(key, testData)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected1 {
		t.Error("Expected ", expected1, " got ", res)
	}

	// test key isn't StringItem
	storage[key] = &ListItem{}
	res, err = Append(key, "")
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != ":0" {
		t.Error("Expected :0, got ", res)
	}

	// test key is StringItem
	si := &StringItem{Data: "Hello"}
	storage[key] = si
	expected2 := fmt.Sprintf(":%d", len(si.Data)+len(testData))

	res, err = Append(key, testData)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected2 {
		t.Error("Expected ", expected2, ", got ", res)
	}
}

func TestGetSet(t *testing.T) {
	key := "getset"
	testData := " World!"

	// test key doesn't exist in storage
	res, err := GetSet(key, testData)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key isn't StringItem
	storage[key] = &ListItem{}
	res, err = GetSet(key, "")
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "$-1" {
		t.Error("Expected $-1, got ", res)
	}

	// test key is StringItem
	oldData := "hello"
	si := &StringItem{Data: oldData}
	storage[key] = si
	expected := fmt.Sprintf("$%d\r\n%s", len(oldData), oldData)

	res, err = GetSet(key, testData)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}

func TestStrlen(t *testing.T) {
	key := "strlen"
	testData := " World!"

	// test key doesn't exist in storage
	res, err := Strlen(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != "-1" {
		t.Error("Expected -1, got ", res)
	}

	// test key isn't StringItem
	storage[key] = &ListItem{}
	res, err = Strlen(key)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != "-1" {
		t.Error("Expected -1, got ", res)
	}

	// test key is StringItem
	si := &StringItem{Data: testData}
	storage[key] = si
	expected := fmt.Sprintf(":%d", len(testData))

	res, err = Strlen(key)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected {
		t.Error("Expected ", expected, ", got ", res)
	}
}

func TestIncrBy(t *testing.T) {
	key := "incrby"
	by := 5

	// test key doesn't exist in storage
	expected1 := fmt.Sprintf(":%d", by)
	res, err := IncrBy(key, by)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected1 {
		t.Error("Expected ", expected1, " got ", res)
	}

	// test key isn't StringItem
	storage[key] = &ListItem{}
	res, err = IncrBy(key, by)
	if err != ErrorWrongType {
		t.Error("Expected ErrorWrongType, got ", err)
	}
	if res != ":0" {
		t.Error("Expected :0, got ", res)
	}

	// test item isn't integer
	storage[key] = &StringItem{Data: "sadasd"}
	res, err = IncrBy(key, by)
	if err != ErrorIsNotInteger {
		t.Error("Expected ErrorIsNotInteger, got ", err)
	}
	if res != ":0" {
		t.Error("Expected :0, got ", res)
	}

	// test item is integer
	siData := 10
	si := &StringItem{Data: strconv.Itoa(siData)}
	storage[key] = si
	expected2 := fmt.Sprintf(":%d", siData+by)

	res, err = IncrBy(key, by)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if res != expected2 {
		t.Error("Expected ", expected2, ", got ", res)
	}
}
