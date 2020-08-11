package models

type Services struct {
	ID          uint  `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Name        string  `gorm:"not null" json:"name" binding:"required"`
	Description string  `gorm:"type:text" json:"description"`
}
