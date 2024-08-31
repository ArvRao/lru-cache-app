package main

import (
	"context"
	"log"
	"lru-cache-app/cache"
	"lru-cache-app/handlers"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}

	// Get cache capacity from environment variable or use default
	cacheCapacity := 1024
	if os.Getenv("CACHE_CAPACITY") != "" {
		cacheCapacity, err = strconv.Atoi(os.Getenv("CACHE_CAPACITY"))
		if err != nil {
			log.Fatalf("Invalid CACHE_CAPACITY: %v", err)
		}
	}

	// Initialize Fiber app
	app := fiber.New()

	// Initialize the LRU cache with the specified capacity
	cache := cache.NewLRUCache(cacheCapacity)

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // You can specify multiple origins separated by commas
		AllowMethods: "GET,POST",              // Allow only GET and POST methods
	}))

	// Register routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to LRU Cache Application")
	})
	app.Get("/cache/:key", handlers.GetCacheValue(cache))
	app.Post("/cache", handlers.SetCacheValue(cache))

	// Start server in a goroutine
	go func() {
		if err := app.Listen(":" + port); err != nil {
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
