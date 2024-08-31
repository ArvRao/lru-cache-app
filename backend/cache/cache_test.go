package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestLRUCache_SetAndGet(t *testing.T) {
	cache := NewLRUCache(3)

	// Set items
	cache.Set("a", "value1", 5*time.Second)
	cache.Set("b", "value2", 5*time.Second)
	cache.Set("c", "value3", 5*time.Second)

	// Get items
	if val, ok := cache.Get("a"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}
	if val, ok := cache.Get("b"); !ok || val != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}
	if val, ok := cache.Get("c"); !ok || val != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}
}

func TestLRUCache_Expiration(t *testing.T) {
	cache := NewLRUCache(3)

	// Set an item with a short TTL
	cache.Set("a", "value1", 1*time.Second)

	// Wait for the item to expire
	time.Sleep(2 * time.Second)

	if _, ok := cache.Get("a"); ok {
		t.Errorf("Expected key 'a' to expire")
	}
}

func TestLRUCache_Eviction(t *testing.T) {
	cache := NewLRUCache(2)

	// Add two items
	cache.Set("a", "value1", 5*time.Second)
	cache.Set("b", "value2", 5*time.Second)

	// Add another item, causing the first one to be evicted
	cache.Set("c", "value3", 5*time.Second)

	if _, ok := cache.Get("a"); ok {
		t.Errorf("Expected key 'a' to be evicted")
	}

	if val, ok := cache.Get("b"); !ok || val != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}
	if val, ok := cache.Get("c"); !ok || val != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}
}

func TestLRUCache_UpdateAndMoveToFront(t *testing.T) {
	cache := NewLRUCache(2)

	// Add items
	cache.Set("a", "value1", 5*time.Second)
	cache.Set("b", "value2", 5*time.Second)

	// Access item 'a' and update its value
	cache.Set("a", "updated_value1", 5*time.Second)

	// Add another item, causing the least recently used item ('b') to be evicted
	cache.Set("c", "value3", 5*time.Second)

	if val, ok := cache.Get("a"); !ok || val != "updated_value1" {
		t.Errorf("Expected updated_value1, got %v", val)
	}

	if _, ok := cache.Get("b"); ok {
		t.Errorf("Expected key 'b' to be evicted")
	}

	if val, ok := cache.Get("c"); !ok || val != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}
}

func TestLRUCache_MaxCapacity(t *testing.T) {
	cache := NewLRUCache(1024)

	// Add 1024 items to the cache
	for i := 0; i < 1024; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Set(key, fmt.Sprintf("value%d", i), 5*time.Second)
	}

	// Check that all 1024 items are present
	for i := 0; i < 1024; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := cache.Get(key); !ok || val != fmt.Sprintf("value%d", i) {
			t.Errorf("Expected %v, got %v", fmt.Sprintf("value%d", i), val)
		}
	}

	// Add one more item, which should trigger the eviction of "key0"
	cache.Set("key1024", "value1024", 5*time.Second)

	// Check that "key0" has been evicted
	if _, ok := cache.Get("key0"); ok {
		t.Errorf("Expected key0 to be evicted")
	}

	// Check that the last added item is present
	if val, ok := cache.Get("key1024"); !ok || val != "value1024" {
		t.Errorf("Expected value1024, got %v", val)
	}

	// Optionally, check that all other keys from "key1" to "key1023" are still present
	for i := 1; i < 1024; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := cache.Get(key); !ok || val != fmt.Sprintf("value%d", i) {
			t.Errorf("Expected %v, got %v", fmt.Sprintf("value%d", i), val)
		}
	}
}
