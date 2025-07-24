package main

import (
	"fmt"
	"github.com/EchidnaTheG/PokeDex/internal"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "1  hello  world  ",
			expected: []string{"1", "hello", "world"},
		}, {
			input:    "                       1 2                      4Dada 211",
			expected: []string{"1", "2", "4Dada", "211"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word isn't Equal Expected Word")
			}
		}
	}

}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := internal.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := internal.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestCacheConcurrency(t *testing.T) {
	cache := internal.NewCache(5 * time.Second)

	// Start multiple goroutines adding/getting data
	for i := 0; i < 10; i++ {
		go func(id int) {
			key := fmt.Sprintf("key%d", id)
			cache.Add(key, []byte(fmt.Sprintf("value%d", id)))
		}(i)
	}

	// Wait a bit and check if all values were added
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		val, ok := cache.Get(key)
		if !ok {
			t.Errorf("Expected to find key %s", key)
		}
		expected := fmt.Sprintf("value%d", i)
		if string(val) != expected {
			t.Errorf("Expected %s, got %s", expected, string(val))
		}
	}
}

/**func TestGetLocation(t *testing.T){
	cases := []struct{
		input string
		output string
	}{

	}
}
**/
