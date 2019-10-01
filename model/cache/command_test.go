package cache

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
)

var notMatchErr = errors.New("do not match")

type testStruct struct {
	Name string
	Sex  string
}

func TestBase(t *testing.T) {
	s := &testStruct{"test_name", "test_sex"}

	key := []byte("test_key")
	value, _ := json.Marshal(s)

	// we should ensure each test case atomic
	defer Delete(key)

	err := Set(key, value)
	if err != nil {
		t.Fatal(err)
	}

	cacheValue, err := Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(cacheValue, value) == false {
		t.Fatal(notMatchErr)
	}

	isStorageCache := Exist(key)
	if isStorageCache == false {
		t.Fatal(notMatchErr)
	}
}
