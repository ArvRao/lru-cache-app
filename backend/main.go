package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Register routes
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
