package handler

import (
	"boss-stock/database"
	"boss-stock/models"

	"github.com/gofiber/fiber/v2"
)

//Tum urunleri listeleyen GetAllCategory() fonksiyonu tanimlaniyor.
func GetAllCategoryes(c *fiber.Ctx) error {

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

//Verilen ID parametresine göre tek ürün listeleyen fonksiyon
func GetCategory(c *fiber.Ctx) error {

	//Gelen parametre id değişkenine alınıyor (string türünde)
	id := c.Params("id")

	//database.go dosyasindan GetCategory() fonksiyonu cagiriliyor id parametresi argüman gönderiliyor.
	result, err := database.GetCategory(id)
	if err != nil {
		//Eğer err nil'e eşit değilse yani err nesnesi doluysa
		//database.go dosyasindan GetAllCategorys() fonksiyonu hata dondururse 500 yantini ve err nesnesini
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

//Verilen ID parametresine göre tek ürün listeleyen fonksiyon
func DeleteCategory(c *fiber.Ctx) error {

	//Gelen parametre id değişkenine alınıyor (string türünde)
	id := c.Params("id")

	//database.go dosyasindan GetCategory() fonksiyonu cagiriliyor id parametresi argüman gönderiliyor.
	result, err := database.DeleteCategory(id)
	if err != nil {
		//Eğer err nil'e eşit değilse yani err nesnesi doluysa
		//database.go dosyasindan GetAllCategory() fonksiyonu hata dondururse 500 yantini ve err nesnesini
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

//Burada SaveCategory() metodu tanimlaniyor
func SaveCategory(c *fiber.Ctx) error {

	//newCategory degiskenine database.go dosyasindaki Product turunden tanimlama yapiliyor.
	newCategory := new(models.Category)
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
