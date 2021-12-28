package models

import "gorm.io/gorm"

//TODO: implemente edilecek
type Category struct {
	gorm.Model
	Category_name string `json:"category_name"`
	Sub_category  uint   `json:"sub_category"`
}
