package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	Manager  string `json:"manager"` //Mağaza yetkilisi
	Tel      uint   `json:"tel"`
	Mail     string `json:"mail"`
	Password string `json:"_"`
}
