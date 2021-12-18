package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//TODO: Edit Product Attributes

//Product veri yapisi olusturuluyor ve attributes belirleniyor. gorm.Model inherit ediliyor ve json tagler ekleniyor
//gorm.Model ana nesnesi ----  ID, CreatedAt, UpdatedAt ve DeletedAt ---- standart alanlarini iceriyor.
type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Detail      string  `json:"detail"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Barcode     int     `json:"barcode"`
	Store_ID    int     `json:"store_id"`
	Caregory_ID int     `json:"category_id"`
	Entry_Price float64 `json:"entry_price"`
}

//InitDatabase() fonksiyonu tanimlaniyor. Cagirildiginda product.db isimli db aciliyor(yoksa olusturuluyor)
//eger problem yoksa Product{} dizisinin adresi gorm'a gonderiliyor AutoMigrate ile DB'ye veriler yaziliyor.
func InitDatabase() error {
	//gorm.Open() iki parametre aliyo ilki sql driver ve dosya ad, ikinci gorm'un Config türü,
	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	//gorm'a Product{} dizisinin verileri gonderiliyor AutoMigrate() ile veriler yaziliyor.
	db.AutoMigrate(&Product{})

	return nil
}

//GetAllProducts() metodu tanimlaniyor parametre almiyor ancak Product turunde bir dizi return ediyor
func GetAllProducts() ([]Product, error) {
	//products isimli Product turunden bir dizi tanimlaniyor
	var products []Product
	//Yine gorm ile database.go aciliyor
	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		return products, err
	}
	//Burada AutoMigrate() yerine Find() metodu calisiyor ve olusturdugumuz products dizisi arguman geciliyor.
	db.Find(&products)

	return products, nil
}

//TODO: Edit Product Attributes
func CreateProduct(name string, detail string, price float64, quantity int, barcode int, store_id int, category_id int, entry_price float64) (Product, error) {

	var newProduct = Product{Name: name, Detail: detail, Price: price, Quantity: quantity, Barcode: barcode, Store_ID: store_id, Caregory_ID: category_id, Entry_Price: entry_price}

	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		return newProduct, err
	}
	db.Create(&newProduct)

	return newProduct, nil
}

//TODO: Add DeleteProducts
//TODO: Update Products
//TODO: Sell Products
