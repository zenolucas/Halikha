package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	app.Static("/", "./public")

    app.Get("/", login)

    log.Fatal(app.Listen(":3000"))
}

func login(c *fiber.Ctx) error {
	return c.SendFile("public/login.html")
}
