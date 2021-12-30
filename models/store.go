package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name     string `json:"name"`
	Address  string `json:"adres"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Logo     string `json:"logo"`
	Manager  string `json:"manager"` //Mağaza yetkilisi
	Tel      uint   `json:"tel"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	IsActive bool   `json:"isactive"`
}
