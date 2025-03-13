package lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLruCache(t *testing.T) {
	t.Run("cache set test reaching capacity", func(t *testing.T) {
		cache := NewLruCache(3)
		cache.Set("1", "one")
		cache.Set("2", "two")
		cache.Set("3", "three")
		cache.Set("4", "four")
		two, err := cache.Get("2")
		require.EqualValues(t, "two", two)
		require.NoError(t, err)
		three, err := cache.Get("3")
		require.EqualValues(t, "three", three)
		require.NoError(t, err)
		four, err := cache.Get("4")
		require.EqualValues(t, "four", four)
		require.NoError(t, err)
		one, err := cache.Get("1")
		require.EqualValues(t, nil, one)
		require.Error(t, err)
	})

	t.Run("cache set overriding the value at key", func(t *testing.T) {
		cache := NewLruCache(3)
		cache.Set("1", "one")
		one, err := cache.Get("1")
		require.EqualValues(t, "one", one)
		require.NoError(t, err)
		cache.Set("1", "one-two")
		oneTwo, err := cache.Get("1")
		require.EqualValues(t, "one-two", oneTwo)
		require.NoError(t, err)
	})

	t.Run("cache set test with a struct type", func(t *testing.T) {
		type testStruct struct {
			Name   string
			Number int
		}
		cache := NewLruCache(3)
		cache.Set("1", testStruct{Name: "one", Number: 1})
		cache.Set("2", testStruct{Name: "two", Number: 2})
		cache.Set("3", testStruct{Name: "three", Number: 3})
		cache.Set("4", testStruct{Name: "four", Number: 4})
		two, err := cache.Get("2")
		require.EqualValues(t, testStruct{Name: "two", Number: 2}, two)
		require.NoError(t, err)
		three, err := cache.Get("3")
		require.EqualValues(t, testStruct{Name: "three", Number: 3}, three)
		require.NoError(t, err)
		four, err := cache.Get("4")
		require.EqualValues(t, testStruct{Name: "four", Number: 4}, four)
		require.NoError(t, err)
		one, err := cache.Get("1")
		require.EqualValues(t, nil, one)
		require.Error(t, err)
		cache.Set("1", testStruct{Name: "one", Number: 1})
		one, err = cache.Get("1")
		require.EqualValues(t, testStruct{Name: "one", Number: 1}, one)
		require.NoError(t, err)
		two, err = cache.Get("2")
		require.EqualValues(t, nil, two)
		require.Error(t, err)
	})
}
