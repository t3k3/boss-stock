package database

import (
	"boss-stock/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//InitDatabase() fonksiyonu tanimlaniyor. Cagirildiginda product.db isimli db aciliyor(yoksa olusturuluyor)
//eger problem yoksa Product{} dizisinin adresi gorm'a gonderiliyor AutoMigrate ile DB'ye veriler yaziliyor.
func InitDatabase() error {
	//gorm.Open() iki parametre aliyo ilki sql driver ve dosya ad, ikinci gorm'un Config türü,
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	//gorm'a model dosyasından Product{} türünün verileri gonderiliyor AutoMigrate() ile veriler yaziliyor.
	db.AutoMigrate(&model.Product{}, &model.Store{}, &model.Repair{}, &model.Category{})

	return nil
}

//GetAllProducts() metodu tanimlaniyor parametre almiyor ancak Product turunde bir dizi return ediyor
func GetAllProducts() ([]model.Product, error) {
	//products isimli Product turunden bir dizi tanimlaniyor
	var products []model.Product
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return products, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz products dizisi arguman geciliyor.
	db.Find(&products)

	return products, nil
}

//TODO: Edit Product Attributes
func CreateProduct(name string, detail string, price float64, quantity int, barcode uint, store_id uint, category_id uint, entry_price float64) (model.Product, error) {

	var newProduct = model.Product{Name: name, Detail: detail, Price: price, Quantity: quantity, Barcode: barcode, Store_ID: store_id, Category_id: category_id, Entry_Price: entry_price}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newProduct, err
	}
	db.Create(&newProduct)

	return newProduct, nil
}

//GetAllRepairs() metodu tanimlaniyor parametre almiyor ancak Repair turunde bir dizi return ediyor
func GetAllRepairs() ([]model.Repair, error) {
	//repairs isimli Repair turunden bir dizi tanimlaniyor
	var repairs []model.Repair
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return repairs, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz repairs dizisi arguman geciliyor.
	db.Find(&repairs)

	return repairs, nil
}

//TODO: Edit Repair Attributes
func CreateRepair(name string, tel uint, problem string, status string, notes string, estimated_price float64, producer string, device_model string, color string, diagnosis string, sms bool) (model.Repair, error) {

	var newRepair = model.Repair{Name: name, Tel: tel, Problem: problem, Status: status, Notes: notes, Estimated_price: estimated_price, Producer: producer, Device_model: device_model, Color: color, Diagnosis: diagnosis, Sms: sms}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newRepair, err
	}
	db.Create(&newRepair)

	return newRepair, nil
}

//GetAllStores() metodu tanimlaniyor parametre almiyor ancak Store turunde bir dizi return ediyor
func GetAllStores() ([]model.Store, error) {
	//stores isimli Store turunden bir dizi tanimlaniyor
	var stores []model.Store
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return stores, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz stores dizisi arguman geciliyor.
	db.Find(&stores)

	return stores, nil
}

//TODO: Edit Store Attributes
func CreateStore(name string, logo string, manager string, tel uint, mail string) (model.Store, error) {

	var newStore = model.Store{Name: name, Logo: logo, Manager: manager, Tel: tel, Mail: mail}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newStore, err
	}
	db.Create(&newStore)

	return newStore, nil
}

//GetAllCategoryes() metodu tanimlaniyor parametre almiyor ancak Store turunde bir dizi return ediyor
func GetAllCategoryes() ([]model.Category, error) {
	//categories isimli Category turunden bir dizi tanimlaniyor
	var categories []model.Category
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return categories, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz categories dizisi arguman geciliyor.
	db.Find(&categories)

	return categories, nil
}

//TODO: Edit Category Attributes
func CreateCategory(category_name string, sub_category uint) (model.Category, error) {

	var newCategory = model.Category{Category_name: category_name, Sub_category: sub_category}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newCategory, err
	}
	db.Create(&newCategory)

	return newCategory, nil
}

//TODO: Add DeleteProducts
//TODO: Update Products
//TODO: Sell Products
