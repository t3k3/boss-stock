package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//TODO: Edit Product Attributes
type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price string `json:"price"`
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
func CreateProduct(name string, price string) (Product, error) {
	var newProduct = Product{Name: name, Price: price}

	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		return newProduct, err
	}
	db.Create(&Product{Name: name, Price: price})

	return newProduct, nil
}

//TODO: Add DeleteProducts
//TODO: Update Products
//TODO: Sell Products
