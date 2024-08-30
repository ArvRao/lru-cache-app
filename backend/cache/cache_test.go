package cache

import (
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
