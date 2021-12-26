package main

//Kullanacagimiz paketler projeye ekleniyor
import (
	//Bu iki paket projemizdeki dizinler ve ayni isimli .go dosyalaridir.
	"boss-stock/controller"
	"boss-stock/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	//Eger /api/v1/repair/new endpointine POST istegi gelirse repair.go icinden SaveRepair() fonksiyonunu response et.
	app.Post("/api/v1/repair/new", controller.SaveRepair)

	//Eger /api/v1/store endpointine GET istegi gelirse store.go icinden GetAllStores() fonksiyonunu response et.
	app.Get("/api/v1/store", controller.GetAllStores)

	//Eger /api/v1/store/new endpointine POST istegi gelirse store.go icinden SaveStore() fonksiyonunu response et.
	app.Post("/api/v1/store/new", controller.SaveStore)

	//Eger /api/v1/category endpointine GET istegi gelirse category.go icinden GetAllCategorys() fonksiyonunu response et.
	app.Get("/api/v1/category", controller.GetAllCategoryes)

	//Eger /api/v1/category/new endpointine POST istegi gelirse category.go icinden SaveCategory() fonksiyonunu response et.
	app.Post("/api/v1/category/new", controller.SaveCategory)

	//Eğer /dashboard endpointe istek gelirse web sayfasından ram, istekler, kullanımlar vs monitör edilebilir.
	app.Get("/dashboard", monitor.New())
}

func main() {
	app := fiber.New()
	//dbErr ile database.go dosyasindan InitDatabase() fonksiyonunu cagir.
	//InitDatabase() fonksiyonu GORM metodlarini cagiriyor database dosyasini aciyor ve migrate yapiyor.
	//Eger dosya acilirsa []Product dizini db'ye migrate ediyor.
	dbErr := database.InitDatabase()

	//Her istekte terminale log yazdırır.
	app.Use(logger.New())

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
