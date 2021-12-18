package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//TODO: Edit Product Attributes
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

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Product{})

	return nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product

	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		return products, err
	}

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
	db.Create(&Product{Name: name, Detail: detail, Price: price, Quantity: quantity, Barcode: barcode, Store_ID: store_id, Caregory_ID: category_id, Entry_Price: entry_price})

	return newProduct, nil
}

//TODO: Add DeleteProducts
//TODO: Update Products
//TODO: Sell Products
