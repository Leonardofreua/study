package cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {

	t.Run("Put number and Get is correct", func(t *testing.T) {
		cache := Init(1)

		key, value := 1, 1
		cache.Put(key, value)
		got := cache.Get(key)

		if got != value {
			t.Errorf("expected: %v, got: %v", value, got)
		}
	})

	t.Run("Get data not stored in cache should return -1", func(t *testing.T) {
		cache := Init(1)

		got := cache.Get(99)

		if got != -1 {
			t.Errorf("expected: -1, got: %v", got)
		}
	})

	t.Run("Put data over capacity and Get should return -1", func(t *testing.T) {
		cache := Init(3)

		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3)
		cache.Put(4, 4) // excess

		got := cache.Get(1)

		if got != -1 {
			t.Errorf("expected: -1, got: %v", got)
		}
	})

	t.Run("Put key over capacity but recent key exists", func(t *testing.T) {
		cache := Init(3)

		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3)
		cache.Put(4, 4) // excess

		got := cache.Get(1)
		if got != -1 {
			t.Errorf("expected: -1, got: %v", got)
		}

		expected := 4
		got = cache.Get(4)

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})
}
