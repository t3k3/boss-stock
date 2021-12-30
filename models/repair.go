package models

import "gorm.io/gorm"

type Repair struct {
	gorm.Model
	Name            string  `json:"name"`
	Tel             uint    `json:"tel"`
	Problem         string  `json:"problem"`         //Cihazdaki belirttigi problemler
	Status          string  `json:"status"`          //Cihazın tamir durumu (HAZIR - BEKLİYOR - TESLİm EDİLDİ)
	Notes           string  `json:"notes"`           //Ek notlar
	Estimated_price float64 `json:"estimated_price"` //Tahmini fiyat (100 - 200 arası veya tek fiyat 200 lira)
	Brand           string  `json:"brand"`           //Cihaz markası
	Device_model    string  `json:"model"`           //Cihaz Modeli
	Color           string  `json:"color"`           //Cihaz rengi
	Diagnosis       string  `json:"diagnosis"`       //Cihaza konan tanılar belirlenmiş arızalar
	Sms             bool    `json:"sms"`             //SMS gonderilsin mi?
}
