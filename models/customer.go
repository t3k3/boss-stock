package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string    `json:"name"`
	Address     string    `json:"adres"`
	City        string    `json:"city"`
	Region      string    `json:"region"`
	Tel         uint      `json:"tel"`
	Mail        string    `json:"mail"`
	IsActive    bool      `json:"isactive"`
	PaymentDate time.Time `json:"payment_date"`
	Balance     float64   `json:"vade"`
}
