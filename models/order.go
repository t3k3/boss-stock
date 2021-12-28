package models

import "gorm.io/gorm"

//TODO: implemente edilecek
type Order struct {
	gorm.Model
	Product_id     string  `json:"product_id"`
	Price          float64 `json:"price"`          //Satış fiyatı
	Quantity       int     `json:"quantity"`       //Satış Adet
	Barcode        uint    `json:"barcode"`        //Barcode numaraları
	Payment_method float64 `json:"payment_method"` //Alış fiyatı
}
