package main

import (
	"context"
	"log"
	"lru-cache-app/cache"
	"lru-cache-app/handlers"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // You can specify multiple origins separated by commas
		AllowMethods: "GET,POST",              // Allow only GET and POST methods
	}))

	// Register routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to LRU Cache Application")
	})
	// Initialize LRU cache with a capacity of 1024
	cache := cache.NewLRUCache(1024)
	app.Get("/cache/:key", handlers.GetCacheValue(cache))
	app.Post("/cache", handlers.SetCacheValue(cache))

	// Start server in a goroutine
	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Set up graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-c

	log.Println("Gracefully shutting down...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited")
}
