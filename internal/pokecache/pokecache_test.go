package pokecache

import (
	"testing"
	"time"
)

func TestCacheGetAdd(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://somelink.com",
			value: []byte("some data"),
		},
	}

	for _, c := range cases {
		cache := NewCache(interval)
		cache.Add(c.key, c.value)
		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key")
		}
		if string(val) != string(c.value) {
			t.Errorf("expected to find value")
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime * 2

	cache := NewCache(baseTime)
	cache.Add("https://somelink.com", []byte("some text"))

	if _, ok := cache.Get("https://somelink.com"); !ok {
		t.Errorf("expected to find key")
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("https://somelink.com"); ok {
		t.Errorf("expected to not find key")
	}
}
