package controller

import (
	"boss-stock/database"
	"boss-stock/model"

	"github.com/gofiber/fiber/v2"
)

//Tum urunleri listeleyen GetAllProducts() fonksiyonu tanimlaniyor.
func GetAllProducts(c *fiber.Ctx) error {

	//database.go dosyasindan GetAllProducts() fonksiyonu cagiriliyor.
	result, err := database.GetAllProducts()
	if err != nil {
		//Eğer err nil'e eşit değilse yani err nesnesi doluysa
		//database.go dosyasindan GetAllProducts() fonksiyonu hata dondururse 500 yantini ve err nesnesini
		//JSON olarak donduruyor
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	//Hata dondurmedigi durumda 200 yanitiyla yine JSON veri donduruluyor(Burada son eklenen urun donuyor)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

//Burada SaveProduct() metodu tanimlaniyor
func SaveProduct(c *fiber.Ctx) error {

	//newProduct degiskenine database.go dosyasindaki Product turunden tanimlama yapiliyor.
	newProduct := new(model.Product)
	//newProduct ile gelen veri BodyParser ie parse ediliyor.
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
	//database.go dosyasindan CreatProduct() metoduna veriler parametre gecilerek cagri yapiliyor.
	result, err := database.CreateProduct(newProduct.Name, newProduct.Detail, newProduct.Price, newProduct.Quantity, newProduct.Barcode, newProduct.StoreID, newProduct.CategoryID, newProduct.Entry_Price)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//database.CreateProduct() basarili olursa success:true donuyor 200 ile birlikte
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
