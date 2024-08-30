package handlers

import (
	"lru-cache-app/cache"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetCacheValue handles GET requests to retrieve a value from the cache by key
func GetCacheValue(cache *cache.LRUCache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")

		if value, ok := cache.Get(key); ok {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"key":   key,
				"value": value,
			})
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "key not found",
		})
	}
}

// SetCacheValue handles POST requests to set a key-value pair in the cache
func SetCacheValue(cache *cache.LRUCache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request struct {
			Key   string        `json:"key"`
			Value interface{}   `json:"value"`
			TTL   time.Duration `json:"ttl"` // Time-to-live in seconds
		}

		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "failed to parse request",
			})
		}

		// Set the key-value pair in the cache with the given TTL
		cache.Set(request.Key, request.Value, time.Second*time.Duration(request.TTL))

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "success",
		})
	}
}
