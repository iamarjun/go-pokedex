package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()

	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

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
