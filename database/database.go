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
	Price       float64 `json:"price"`       //Satış fiyatı
	Quantity    int     `json:"quantity"`    //Stok Adet
	Barcode     uint    `json:"barcode"`     //Barcode numarası
	Store_ID    uint    `json:"store_id"`    //Mağaza ID, ürünün hangi mağazaya ait olduğu (BAKILACAK)
	Caregory_ID uint    `json:"category_id"` //Ürünün kaegorisi
	Entry_Price float64 `json:"entry_price"` //Alış fiyatı
}

type Store struct {
	gorm.Model
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	Manager string `json:"manager"` //Mağaza yetkilisi
	Tel     uint   `json:"tel"`
	Mail    string `json:"mail"`
}

type Repair struct {
	gorm.Model
	Name            string  `json:"name"`
	Tel             uint    `json:"tel"`
	Problem         string  `json:"problem"`         //Cihazdaki belirttigi problemler
	Status          string  `json:"status"`          //Cihazın tamir durumu (HAZIR - BEKLİYOR - TESLİm EDİLDİ)
	Notes           string  `json:"notes"`           //Ek notlar
	Estimated_price float64 `json:"estimated_price"` //Tahmini fiyat (100 - 200 arası veya tek fiyat 200 lira)
	Producer        string  `json:"producer"`        //Cihaz markası
	Device_model    string  `json:"model"`           //Cihaz Modeli
	Color           string  `json:"color"`           //Cihaz rengi
	Diagnosis       string  `json:"diagnosis"`       //Cihaza konan tanılar belirlenmiş arızalar
	Sms             bool    `json:"sms"`             //SMS gonderilsin mi?
}

//InitDatabase() fonksiyonu tanimlaniyor. Cagirildiginda product.db isimli db aciliyor(yoksa olusturuluyor)
//eger problem yoksa Product{} dizisinin adresi gorm'a gonderiliyor AutoMigrate ile DB'ye veriler yaziliyor.
func InitDatabase() error {
	//gorm.Open() iki parametre aliyo ilki sql driver ve dosya ad, ikinci gorm'un Config türü,
	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	//gorm'a Product{} dizisinin verileri gonderiliyor AutoMigrate() ile veriler yaziliyor.
	db.AutoMigrate(&Product{}, &Store{}, &Repair{})

	return nil
}

//GetAllProducts() metodu tanimlaniyor parametre almiyor ancak Product turunde bir dizi return ediyor
func GetAllProducts() ([]Product, error) {
	//products isimli Product turunden bir dizi tanimlaniyor
	var products []Product
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
func CreateProduct(name string, detail string, price float64, quantity int, barcode uint, store_id uint, category_id uint, entry_price float64) (Product, error) {

	var newProduct = Product{Name: name, Detail: detail, Price: price, Quantity: quantity, Barcode: barcode, Store_ID: store_id, Caregory_ID: category_id, Entry_Price: entry_price}

	db, err := gorm.Open(sqlite.Open("bossdb.db"), &gorm.Config{})
	if err != nil {
		return newProduct, err
	}
	db.Create(&newProduct)

	return newProduct, nil
}

//GetAllRepairs() metodu tanimlaniyor parametre almiyor ancak Repair turunde bir dizi return ediyor
func GetAllRepairs() ([]Repair, error) {
	//repairs isimli Repair turunden bir dizi tanimlaniyor
	var repairs []Repair
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
func CreateRepair(name string, tel uint, problem string, status string, notes string, estimated_price float64, producer string, device_model string, color string, diagnosis string, sms bool) (Repair, error) {

	var newRepair = Repair{Name: name, Tel: tel, Problem: problem, Status: status, Notes: notes, Estimated_price: estimated_price, Producer: producer, Device_model: device_model, Color: color, Diagnosis: diagnosis, Sms: sms}

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
