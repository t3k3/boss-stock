package controller

import (
	"boss-stock/database"
	"boss-stock/model"

	"github.com/gofiber/fiber/v2"
)

//Tum urunleri listeleyen GetAllStores() fonksiyonu tanimlaniyor.
func GetAllStores(c *fiber.Ctx) error {
	//database.go dosyasindan GetAllStores() fonksiyonu cagiriliyor.
	result, err := database.GetAllStores()
	if err != nil {
		//database.go dosyasindan GetAllStores() fonksiyonu hata dondururse 500 yantini JSON olarak donduruyor
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	//Hata dondurmedigi durumda 200 yanitiyla yine JSON veri donduruluyor(Burada son eklenen store donuyor)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

//Burada SaveStores() metodu tanimlaniyor
func SaveStore(c *fiber.Ctx) error {
	//newStore degiskenine database.go dosyasindaki Product turunden tanimlama yapiliyor.
	newStore := new(model.Store)
	//newStore ile gelen veri BodyParser ie parse ediliyor.
	err := c.BodyParser(newStore)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//TODO: Edit Store Attributes
	//database.go dosyasindan CreateStore() metoduna veriler parametre gecilerek cagri yapiliyor.
	result, err := database.CreateStore(newStore.Name, newStore.Logo, newStore.Manager, newStore.Tel, newStore.Mail)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//database.CreateStore() basarili olursa success:true donuyor 200 ile birlikte
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
