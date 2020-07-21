package models

type Category struct {
	ID   uint   `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}
