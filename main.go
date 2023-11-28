package main

import (
	"caesar-cipher/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create the HTML template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber app with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files from the "static" directory
	app.Static("/static", "./static")

	// Initialize routes
	routes.RouteInit(app)

	// Start the Fiber app
	log.Fatal(app.Listen(":8000"))
}
