package models

import "gorm.io/gorm"

//TODO: implemente edilecek
type Order struct {
	gorm.Model
	OrderProducts string  `json:"orders"`         //Siparişteki ürünler Product türünden liste şeklinde tutulacak
	TotalPrice    float64 `json:"total_price"`    //Toplam fiyatı
	RealPrice     float64 `json:"real_price"`     //indirim sonrası toplam fiyat
	PaymentMethod uint    `json:"payment_method"` //Ödeme metodu (1 = veresiye 2 = nakit 3 = kart 4 = online)
	SalesType     uint    `json:"sales_type"`     //Satış türünü belirtir. (1 = mağazadan 2 = online)
}
