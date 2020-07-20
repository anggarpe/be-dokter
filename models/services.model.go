package models

type Services struct {
	ID       string  `gorm:"primary_key;unique;not null" json:"id"`
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Price float64 `gorm:"type:double" json:"price"`
	Description string `gorm:"type:text" json:"description"`
}