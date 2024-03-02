package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error { // pass by reference

	// return json 
	return c.JSON(fiber.Map{
		"message": "Hello, World 👋!",
	})
	
	// return c.SendString("Hello, World 👋!")
}