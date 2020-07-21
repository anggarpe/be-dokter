package models

type Admin struct {
	ID       string  `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	Password string `json:"password"`
}

//type Admin []Admin
