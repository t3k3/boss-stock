package models

import "gorm.io/gorm"

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
	StoreID     uint    `json:"store_id"`    //StoreID Foreign Key ???
	CategoryID  uint    `json:"category_id"` //Ürünün kaegorisi
	Entry_Price float64 `json:"entry_price"` //Alış fiyatı
}
