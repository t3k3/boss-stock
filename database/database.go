package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
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

func CreateProduct(name string, url string) (Product, error) {
	var newProduct = Product{Name: name, Url: url}

	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		return newProduct, err
	}
	db.Create(&Product{Name: name, Url: url})

	return newProduct, nil
}
