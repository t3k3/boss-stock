package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!" + c.Params("name"))
	})

	app.Listen(":3000")
}
