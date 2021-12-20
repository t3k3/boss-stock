package model

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
