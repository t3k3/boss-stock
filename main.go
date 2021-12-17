package main

import (
	"bookmark-api-fiber/database"
	"bookmark-api-fiber/product"

	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/v1/product", product.GetAllProducts)
	app.Post("/api/v1/product", product.SaveProduct)
}

func main() {
	app := fiber.New()
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
