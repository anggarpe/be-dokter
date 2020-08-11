package models

import (
	"time"
)

type Order struct {
	ID string `gorm:"type:int(5);primary_key;unique; not null" json:"id"`
	Total float64 `gorm:"type:double" json:"total"`
	ShippedAddress string `gorm:"type:text" json:"shippedAddress"`
	ShippedDate *time.Time `json:"shippedDate"`
	Status string `json:"status"`
	OrderDate *time.Time `json:"orderDate"`
	UserID uint `sql:"index" json:"-"`
	User User `json:"user"`
	OrderDetail []OrderDetail
}
