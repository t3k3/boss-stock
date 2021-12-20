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
	db.AutoMigrate(&model.Product{}, &model.Store{}, &model.Repair{})

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

	var newProduct = model.Product{Name: name, Detail: detail, Price: price, Quantity: quantity, Barcode: barcode, Store_ID: store_id, Caregory_ID: category_id, Entry_Price: entry_price}

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

//TODO: Add DeleteProducts
//TODO: Update Products
//TODO: Sell Products
