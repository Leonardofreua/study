package cache

import (
	"testing"
)

func TestLFUCache(t *testing.T) {

	t.Run("Put number and Get is correct", func(t *testing.T) {
		cache := InitLFU(1)

		key, value := 1, 1
		cache.Put(key, value)
		got := cache.Get(key)

		if got != value {
			t.Errorf("expected: %v, got: %v", value, got)
		}
	})

	t.Run("Get data not stored in cache should return nil", func(t *testing.T) {
		cache := InitLFU(1)
		got := cache.Get(99)

		if got != nil {
			t.Errorf("expected nil, got: %v", got)
		}
	})

	t.Run("Should update the item's value and frequency when it exists", func(t *testing.T) {
		cache := InitLFU(1)

		key, value1 := 1, 1
		cache.Put(key, value1)

		if cache.MinFreq != 1 {
			t.Errorf("expected MinFreq: 1, got: %v ", cache.MinFreq)
		}

		if len(cache.Frequency) != 1 {
			t.Errorf("expected Frequency: 1, got: %v ", len(cache.Frequency))
		}

		if cache.Items[key].Value.(*LFUNode).Frequency != 1 {
			t.Errorf("expected node Frequency: 1, got: %v ", cache.Items[key].Value.(*LFUNode).Frequency)
		}

		got := cache.Get(key)

		if len(cache.Frequency) != 2 {
			t.Errorf("expected Frequency: 2, got: %v ", len(cache.Frequency))
		}

		if cache.Items[key].Value.(*LFUNode).Frequency != 2 {
			t.Errorf("expected node Frequency: 2, got: %v ", cache.Items[key].Value.(*LFUNode).Frequency)
		}

		if got != value1 {
			t.Errorf("expected: %v, got: %v", value1, got)
		}

		value2 := 2
		cache.Put(key, value2)

		if len(cache.Frequency) != 3 {
			t.Errorf("expected Frequency: 3, got: %v ", len(cache.Frequency))
		}

		if cache.Items[key].Value.(*LFUNode).Frequency != 3 {
			t.Errorf("expected node Frequency: 3, got: %v ", cache.Items[key].Value.(*LFUNode).Frequency)
		}

		got = cache.Get(key)

		if len(cache.Frequency) != 4 {
			t.Errorf("expected Frequency: 4, got: %v ", len(cache.Frequency))
		}

		if cache.Items[key].Value.(*LFUNode).Frequency != 4 {
			t.Errorf("expected Frequency: 4, got: %v ", cache.Items[key].Value.(*LFUNode).Frequency)
		}

		if got != value2 {
			t.Errorf("expected: %v, got: %v", value2, got)
		}
	})
}

func TestPuttingDataOverCapacityGetShouldReturnNil(t *testing.T) {
	var tests = []struct {
		capacity   int
		operations map[int]int
		getKey     int
	}{
		{
			0,
			map[int]int{1: 1},
			1,
		},
		{
			2,
			map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
			1,
		},
	}

	for _, tt := range tests {
		t.Run("Put data over capacity and Get should return nil", func(t *testing.T) {
			cache := InitLFU(tt.capacity)

			for key, value := range tt.operations {
				cache.Put(key, value)
			}

			got := cache.Get(tt.getKey)
			if got != nil {
				t.Errorf("Expected nil, got: %v", got)
			}
		})
	}
}
