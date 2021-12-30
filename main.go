package main

//Kullanacagimiz paketler projeye ekleniyor
import (
	//Bu iki paket projemizdeki dizinler ve ayni isimli .go dosyalaridir.
	"boss-stock/database"
	"boss-stock/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

//status isimli bu fonksiyon SendString() ile girdigimiz string degeri donduruyor.
//Bu ve benzeri fonksiyonları gelen endpointlere göre parametre geçiyoruz.
func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

//setupRoutes fonksiyonunu implemente ediyoruz, icinde route'larimizi tanimliyoruz.
func setupRoutes(app *fiber.App) {

	//Handlers

	app.Get("/", status)

	app.Get("/api/v1/product", handler.GetAllProducts)

	app.Get("/api/v1/product/:id", handler.GetProduct)

	app.Get("/api/v1/product/category/:id", handler.GetProductByCategoryID)

	app.Get("/api/v1/product/barcode/:id", handler.GetProductByBarcode)

	app.Delete("/api/v1/product/:id", handler.DeleteProduct)

	app.Post("/api/v1/product/new", handler.SaveProduct)

	app.Get("/api/v1/repair", handler.GetAllRepairs)

	app.Get("/api/v1/repair/:id", handler.GetRepair)

	app.Delete("/api/v1/repair/:id", handler.DeleteRepair)

	app.Post("/api/v1/repair/new", handler.SaveRepair)

	app.Get("/api/v1/store", handler.GetAllStores)

	app.Get("/api/v1/store/:id", handler.GetStore)

	app.Delete("/api/v1/store/:id", handler.DeleteStore)

	app.Post("/api/v1/store/new", handler.SaveStore)

	app.Get("/api/v1/category", handler.GetAllCategoryes)

	app.Get("/api/v1/category/:id", handler.GetCategory)

	app.Delete("/api/v1/category/:id", handler.DeleteCategory)

	app.Post("/api/v1/category/new", handler.SaveCategory)

	//Eğer /dashboard endpointe istek gelirse web sayfasından ram, istekler, kullanımlar vs monitör edilebilir.
	app.Get("/dashboard", monitor.New())
}

func main() {
	app := fiber.New()
	//dbErr ile database.go dosyasindan InitDatabase() fonksiyonunu cagir. InitDatabase() error
	//nesnesi döndürür.
	//InitDatabase() fonksiyonu GORM metodlarini cagiriyor database dosyasini aciyor ve migrate yapiyor.
	//Eger dosya acilirsa database içinde tanımladığımız InitDatabase() fonksiyonu çalışacak.
	//InitDatabase içinde AutoMigrate() metoduna verdiğimiz parametlere models içindeki gibi
	//database tablolarını oluşturacak. (&models.Product{}, &models.Store{}, &models.Repair{}, &models.Category{})
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	//Her istekte terminale log yazdırır.
	app.Use(logger.New())

	setupRoutes(app)
	app.Listen(":3000")
}
