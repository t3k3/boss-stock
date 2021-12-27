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
//Bu ve benzeri fonksiyonları gelen endpointlere göre parametre geçiyoruz.
func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

//setupRoutes fonksiyonunu implemente ediyoruz, icinde route'larimizi tanimliyoruz.
func setupRoutes(app *fiber.App) {

	//Eger / dzinine istek gelirse status fonksiyonunu response et.
	app.Get("/", status)

	app.Get("/api/v1/product", controller.GetAllProducts)

	app.Post("/api/v1/product/new", controller.SaveProduct)

	app.Get("/api/v1/repair", controller.GetAllRepairs)

	app.Post("/api/v1/repair/new", controller.SaveRepair)

	app.Get("/api/v1/store", controller.GetAllStores)

	app.Post("/api/v1/store/new", controller.SaveStore)

	app.Get("/api/v1/category", controller.GetAllCategoryes)

	app.Post("/api/v1/category/new", controller.SaveCategory)

	//Eğer /dashboard endpointe istek gelirse web sayfasından ram, istekler, kullanımlar vs monitör edilebilir.
	app.Get("/dashboard", monitor.New())
}

func main() {
	app := fiber.New()
	//dbErr ile database.go dosyasindan InitDatabase() fonksiyonunu cagir. InitDatabase() error
	//nesnesi döndürür.
	//InitDatabase() fonksiyonu GORM metodlarini cagiriyor database dosyasini aciyor ve migrate yapiyor.
	//Eger dosya acilirsa database içinde tanımladığımız InitDatabase() fonksiyonu çalışacak.
	//InitDatabase içinde AutoMigrate() metoduna verdiğimiz parametlere model içindeki gibi
	//database tablolarını oluşturacak. (&model.Product{}, &model.Store{}, &model.Repair{}, &model.Category{})
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	//Her istekte terminale log yazdırır.
	app.Use(logger.New())

	setupRoutes(app)
	app.Listen(":3000")
}
