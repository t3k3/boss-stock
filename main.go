package main

//Kullanacagimiz paketler projeye ekleniyor
import (
	//Bu iki paket projemizdeki dizinler ve ayni isimli .go dosyalaridir.
	"boss-stock/controller"
	"boss-stock/database"

	"github.com/gofiber/fiber/v2"
)

//status isimli bu fonksiyon SendString() ile girdigimiz string degeri donduruyor.
func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

//setupRoutes fonksiyonu icinde route'larimizi tanimliyoruz.
func setupRoutes(app *fiber.App) {

	//Eger / dzinine istek gelirse status fonksiyonunu response et.
	app.Get("/", status)
	//TODO: Edit Endpoints

	//Eger /api/v1/product endpointine GET istegi gelirse product.go icinden GetAllProducts() fonksiyonunu response et.
	app.Get("/api/v1/product", controller.GetAllProducts)
	//Eger /api/v1/product/new endpointine POST istegi gelirse product.go icinden SaveProduct() fonksiyonunu response et.
	app.Post("/api/v1/product/new", controller.SaveProduct)

	//Eger /api/v1/repair endpointine GET istegi gelirse repair.go icinden GetAllRepairs() fonksiyonunu response et.
	app.Get("/api/v1/repair", controller.GetAllRepairs)
	//Eger /api/v1/repair endpointine POST istegi gelirse product.go icinden SaveRepair() fonksiyonunu response et.
	app.Post("/api/v1/repair/new", controller.SaveRepair)
}

func main() {
	app := fiber.New()
	//dbErr ile database.go dosyasindan InitDatabase() fonksiyonunu cagir.
	//InitDatabase() fonksiyonu GORM metodlarini cagiriyor database dosyasini aciyor ve migrate yapiyor.
	//Eger dosya acilirsa []Product dizini db'ye migrate ediyor.
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
