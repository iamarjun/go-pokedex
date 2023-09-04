package pokecache

import (
	"testing"
	"time"
)

const interval = time.Microsecond * 10

func TestCreateCache(t *testing.T) {
	cache := NewCache(interval)

	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(interval)

	cases := []struct {
		inputKey   string
		inputValue []byte
	}{
		{
			inputKey:   "key1",
			inputValue: []byte("val1"),
		},
		{
			inputKey:   "key2",
			inputValue: []byte("val2"),
		},
		{
			inputKey:   "",
			inputValue: []byte("val3"),
		},
	}

	for _, c := range cases {

		cache.Add(c.inputKey, c.inputValue)

		actual, ok := cache.Get(c.inputKey)

		if !ok {
			t.Errorf("%s is not found", c.inputKey)
			continue
		}

		if string(actual) != string(c.inputValue) {
			t.Errorf("%s doesn't match %s", actual, c.inputValue)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(interval)

	key1 := "key1"
	cache.Add(key1, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key1)

	if ok {
		t.Errorf("%s could have been reaped", key1)
	}

}
func TestReapFail(t *testing.T) {
	cache := NewCache(interval)

	key1 := "key1"
	cache.Add(key1, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(key1)

	if ok {
		t.Errorf("%s could have been reaped", key1)
	}

}
