package models

type OrderDetail struct {
	ID uint `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Price float64 `gorm:"type:double" json:"price"`
	Quantity int32 `gorm:"type:int" json:"quantity"`
	OrderID uint `sql:"index" json:"id_order"`
	ProductID uint `sql:"index" json:"id_product"`
}
