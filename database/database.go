package database

import (
	"boss-stock/models"

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
	db.AutoMigrate(&models.Product{}, &models.Store{}, &models.Repair{}, &models.Category{})

	return nil
}

//GetAllProducts() metodu tanimlaniyor parametre almiyor ancak Product turunde bir dizi return ediyor
func GetAllProducts() ([]models.Product, error) {
	//products isimli Product turunden bir dizi tanimlaniyor
	var products []models.Product
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return products, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz products dizisi arguman geciliyor.
	db.Find(&products)

	return products, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func GetProduct(id string) ([]models.Product, error) {
	//product isimli Product turunden bir instance tanimlaniyor
	var product []models.Product
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return product, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz product dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Find(&product, id)

	return product, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func GetProductByCategoryID(id string) ([]models.Product, error) {
	//product isimli Product turunden bir instance tanimlaniyor
	var product []models.Product
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return product, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz product dizisi arguman geciliyor.
	//db.Where() kullanılıyor. Bu şekilde id değeriyle db içindeki category_id karşılaştırılıyor.

	//category_id == id olan verileri getir.
	db.Where("category_id == ?", id).Find(&product)
	return product, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func GetProductByBarcode(id string) ([]models.Product, error) {
	//product isimli Product turunden bir instance tanimlaniyor
	var product []models.Product
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return product, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz product dizisi arguman geciliyor.
	//db.Where() kullanılıyor. Bu şekilde id değeriyle db içindeki category_id karşılaştırılıyor.

	//category_id == id olan verileri getir.
	db.Where("barcode == ?", id).Find(&product)
	return product, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func DeleteProduct(id string) ([]models.Product, error) {
	//product isimli Product turunden bir instance tanimlaniyor
	var product []models.Product
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return product, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz product dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Delete(&product, id)

	return product, nil
}

//TODO: Edit Product Attributes
func CreateProduct(name string, detail string, price float64, quantity int, barcode uint, store_id uint, category_id uint, entry_price float64, kdv float64) (models.Product, error) {

	var newProduct = models.Product{Name: name, Detail: detail, Price: price, Quantity: quantity, Barcode: barcode, StoreID: store_id, CategoryID: category_id, Entry_Price: entry_price, Tax: kdv}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newProduct, err
	}
	db.Create(&newProduct)

	return newProduct, nil
}

//GetAllRepairs() metodu tanimlaniyor parametre almiyor ancak Repair turunde bir dizi return ediyor
func GetAllRepairs() ([]models.Repair, error) {
	//repairs isimli Repair turunden bir dizi tanimlaniyor
	var repairs []models.Repair
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return repairs, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz repairs dizisi arguman geciliyor.
	db.Find(&repairs)

	return repairs, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func GetRepair(id string) ([]models.Repair, error) {
	//repair isimli Repair turunden bir instance tanimlaniyor
	var repair []models.Repair
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return repair, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz repair dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Find(&repair, id)

	return repair, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func DeleteRepair(id string) ([]models.Repair, error) {
	//repair isimli Repair turunden bir instance tanimlaniyor
	var repair []models.Repair
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return repair, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz repair dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Delete(&repair, id)

	return repair, nil
}

//TODO: Edit Repair Attributes
func CreateRepair(name string, tel uint, problem string, status string, notes string, estimated_price float64, brand string, device_model string, color string, diagnosis string, sms bool) (models.Repair, error) {

	var newRepair = models.Repair{Name: name, Tel: tel, Problem: problem, Status: status, Notes: notes, Estimated_price: estimated_price, Brand: brand, Device_model: device_model, Color: color, Diagnosis: diagnosis, Sms: sms}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newRepair, err
	}
	db.Create(&newRepair)

	return newRepair, nil
}

//GetAllStores() metodu tanimlaniyor parametre almiyor ancak Store turunde bir dizi return ediyor
func GetAllStores() ([]models.Store, error) {
	//stores isimli Store turunden bir dizi tanimlaniyor
	var stores []models.Store
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return stores, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz stores dizisi arguman geciliyor.
	db.Find(&stores)

	return stores, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func GetStore(id string) ([]models.Store, error) {
	//store isimli Store turunden bir instance tanimlaniyor
	var store []models.Store
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return store, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz store dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Find(&store, id)

	return store, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func DeleteStore(id string) ([]models.Store, error) {
	//store isimli Store turunden bir instance tanimlaniyor
	var store []models.Store
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return store, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz store dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Delete(&store, id)

	return store, nil
}

//TODO: Edit Store Attributes
func CreateStore(name string, address string, city string, region string, logo string, manager string, tel uint, mail string, password string, isactive bool) (models.Store, error) {

	var newStore = models.Store{Name: name, Address: address, City: city, Region: region, Logo: logo, Manager: manager, Tel: tel, Mail: mail, Password: password, IsActive: isactive}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newStore, err
	}
	db.Create(&newStore)

	return newStore, nil
}

//GetAllCategoryes() metodu tanimlaniyor parametre almiyor ancak Store turunde bir dizi return ediyor
func GetAllCategoryes() ([]models.Category, error) {
	//categories isimli Category turunden bir dizi tanimlaniyor
	var categories []models.Category
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return categories, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz categories dizisi arguman geciliyor.
	db.Find(&categories)

	return categories, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func GetCategory(id string) ([]models.Category, error) {
	//category isimli Category turunden bir instance tanimlaniyor
	var category []models.Category
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return category, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz category dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Find(&category, id)

	return category, nil
}

//Yalnızca bir ürün çeken fonksiyona id bilgisini geçiyoruz.
func DeleteCategory(id string) ([]models.Category, error) {
	//category isimli Category turunden bir instance tanimlaniyor
	var category []models.Category
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return category, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz category dizisi arguman geciliyor.
	//ikinci parametre olarak id veriliyor. Bu şekilde id değeriyle db içindeki id karşılaştırılıyor.
	db.Delete(&category, id)

	return category, nil
}

//TODO: Edit Category Attributes
func CreateCategory(category_name string, sub_category uint) (models.Category, error) {

	var newCategory = models.Category{Category_name: category_name, Sub_category: sub_category}

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
