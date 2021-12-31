package handler

import (
	"boss-stock/database"
	"boss-stock/models"

	"github.com/gofiber/fiber/v2"
)

//Tum urunleri listeleyen GetAllOrders() fonksiyonu tanimlaniyor.
func GetAllOrders(c *fiber.Ctx) error {

	//database.go dosyasindan GetAllOrders() fonksiyonu cagiriliyor.
	result, err := database.GetAllOrders()
	if err != nil {
		//database.go dosyasindan GetAllOrders() fonksiyonu hata dondururse 500 yantini JSON olarak donduruyor
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	//Hata dondurmedigi durumda 200 yanitiyla yine JSON veri donduruluyor(Burada son eklenen category donuyor)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

//Burada SaveOrder() metodu tanimlaniyor
func SaveOrder(c *fiber.Ctx) error {

	//newOrder degiskenine database.go dosyasindaki Product turunden tanimlama yapiliyor.
	newOrder := new(models.Order)
	//newOrder ile gelen veri BodyParser ie parse ediliyor.
	err := c.BodyParser(newOrder)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//TODO: Edit Order Attributes
	//database.go dosyasindan CreateOrder() metoduna veriler parametre gecilerek cagri yapiliyor.
	result, err := database.CreateOrder(newOrder.OrderProducts, newOrder.TotalPrice, newOrder.RealPrice, newOrder.PaymentMethod, newOrder.SalesType)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//database.CreateCategory() basarili olursa success:true donuyor 200 ile birlikte
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
