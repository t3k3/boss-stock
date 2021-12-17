package product

import (
	"boss-stock/database"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	result, err := database.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

func SaveProduct(c *fiber.Ctx) error {
	newProduct := new(database.Product)

	err := c.BodyParser(newProduct)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//TODO: Edit Product Attributes
	result, err := database.CreateProduct(newProduct.Name, newProduct.Price)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
