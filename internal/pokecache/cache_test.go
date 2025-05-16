package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	fmt.Println("Begin addget test")
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "test",
			val: []byte("testval"),
		},
		{
			key: "https://example.com/",
			val: []byte("example website info"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			pokeCache := NewCache(5 * time.Minute)
			pokeCache.Add(c.key, c.val)
			val, ok := pokeCache.Get(cases[i].key)
			if !ok {
				t.Errorf("didn't find key")
				return
			}
			if string(val) != string(cases[i].val) {
				t.Errorf("vals not the same")
				return
			}
		})
	}
	fmt.Println("End addget test")
}

func TestReapLoop(t *testing.T) {
	fmt.Println("Begin reaploop test")
	const (
		interval   = 5 * time.Millisecond
		timeToWait = 10 * time.Millisecond
	)
	pokeCache := NewCache(interval)
	pokeCache.Add("test", []byte("testval"))
	_, ok := pokeCache.Get("test")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(timeToWait)

	_, ok = pokeCache.Get("test")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	fmt.Println("End reaploop test")
	return
}
