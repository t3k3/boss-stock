package controller

import (
	"boss-stock/database"

	"github.com/gofiber/fiber/v2"
)

//Tum urunleri listeleyen GetAllRepairs() fonksiyonu tanimlaniyor.
func GetAllRepairs(c *fiber.Ctx) error {
	//database.go dosyasindan GetAllRepairs() fonksiyonu cagiriliyor.
	result, err := database.GetAllRepairs()
	if err != nil {
		//database.go dosyasindan GetAllRepairs() fonksiyonu hata dondururse 500 yantini JSON olarak donduruyor
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	//Hata dondurmedigi durumda 200 yanitiyla yine JSON veri donduruluyor(Burada son eklenen repair donuyor)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

//Burada SaveRepair() metodu tanimlaniyor
func SaveRepair(c *fiber.Ctx) error {
	//newRepair degiskenine database.go dosyasindaki Product turunden tanimlama yapiliyor.
	newRepair := new(database.Repair)
	//newRepair ile gelen veri BodyParser ie parse ediliyor.
	err := c.BodyParser(newRepair)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//TODO: Edit Repair Attributes
	//database.go dosyasindan CreateRepair() metoduna veriler parametre gecilerek cagri yapiliyor.
	result, err := database.CreateRepair(newRepair.Name, newRepair.Tel, newRepair.Problem, newRepair.Status, newRepair.Notes, newRepair.Estimated_price, newRepair.Producer, newRepair.Device_model, newRepair.Color, newRepair.Diagnosis, newRepair.Sms)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//database.CreateRepair() basarili olursa success:true donuyor 200 ile birlikte
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
