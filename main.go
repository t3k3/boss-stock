package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Sarj Aleti", Price: "120,00 TL"},
	{ID: 2, Name: "Kilif", Price: "75,00 TL"},
}

func getProduct(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(products)
}

func main() {

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Boss Stock!")
	})
	app.Get("/api", getProduct)

	app.Listen(":3001")
}
