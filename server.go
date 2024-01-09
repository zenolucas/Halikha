package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Serve Bootstrap CSS file
	app.Static("/bootstrap/css", "./node_modules/bootstrap/dist/css")

	// Serve Bootstrap JS files
	app.Static("/bootstrap/js", "./node_modules/bootstrap/dist/js")
	app.Static("/bootstrap/js/popper", "./node_modules/@popperjs/core/dist/umd")

	// Serve other static files
	app.Static("/", "./public")

	app.Get("/", login)

	// Authentication route
	app.Post("/authenticate", func(c *fiber.Ctx) error {
		// Retrieve form data
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Simple authentication (replace with your actual authentication logic)
		if username == "user" && password == "password" {
			return c.SendString("Login successful!")
		} else {
			return c.SendString("Login failed. Invalid username or password.")
		}
	})

	log.Fatal(app.Listen(":3000"))
}

func login(c *fiber.Ctx) error {
	return c.SendFile("public/index.html")
}
