package main

import (
	"lru-cache-app/cache"
	"lru-cache-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize LRU cache with a capacity of 1024
	cache := cache.NewLRUCache(1024)

	// Register routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to LRU Cache Application. Hope you'll enjoy it!")
	})
	app.Get("/cache/:key", handlers.GetCacheValue(cache))
	app.Post("/cache", handlers.SetCacheValue(cache))

	// Start the server on port 3000
	app.Listen(":3000")
}
