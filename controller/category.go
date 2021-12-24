package controller

import (
	"boss-stock/database"
	"boss-stock/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//Tum urunleri listeleyen GetAllCategory() fonksiyonu tanimlaniyor.
func GetAllCategoryes(c *fiber.Ctx) error {

	//TODO: TEMP:Gelen istekleri terminalden takip etmek için geçici bir kod
	fmt.Println(c.IP(), "\tGet All Category\t /api/v1/category")

	//database.go dosyasindan GetAllCategory() fonksiyonu cagiriliyor.
	result, err := database.GetAllCategoryes()
	if err != nil {
		//database.go dosyasindan GetAllCategory() fonksiyonu hata dondururse 500 yantini JSON olarak donduruyor
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

//Burada SaveCategory() metodu tanimlaniyor
func SaveCategory(c *fiber.Ctx) error {

	//TODO: TEMP:Gelen istekleri terminalden takip etmek için geçici bir kod
	fmt.Println(c.IP(), "\tSave New Category\t /api/v1/category/new")

	//newCategory degiskenine database.go dosyasindaki Product turunden tanimlama yapiliyor.
	newCategory := new(model.Category)
	//newCategory ile gelen veri BodyParser ie parse ediliyor.
	err := c.BodyParser(newCategory)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	//TODO: Edit Category Attributes
	//database.go dosyasindan CreateCategory() metoduna veriler parametre gecilerek cagri yapiliyor.
	result, err := database.CreateCategory(newCategory.Category_name, newCategory.Sub_category)
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
