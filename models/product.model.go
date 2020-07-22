package models

type Product struct {
	ID uint `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Price float64 `gorm:"type:double" json:"price"`
	Stock int32 `gorm:"type:int" json:"stock"`
	Description string `gorm:"type:text" json:"description"`
	CategoryID uint `sql:"index" json:"-"`
	Category Category `json:"category"`
	OrderDetail []OrderDetail `json:"-"`
}
