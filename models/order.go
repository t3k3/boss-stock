package models

import "gorm.io/gorm"

//TODO: implemente edilecek
type Order struct {
	gorm.Model
	Products      []Product `json:"producs"`        //Siparişteki ürünler Product türünden liste şeklinde tutulacak
	TotalPrice    float64   `json:"total_price"`    //Toplam fiyatı
	PaymentMethod uint      `json:"payment_method"` //Ödeme metodu (1 = veresiye 2 = nakit 3 = kart 4 = online)
	Type          uint      `json:"type"`           //Satış türünü belirtir. (1 = mağazadan 2 = online)
}
