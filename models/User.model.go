package models

type User struct {
	ID       int    `gorm:"primary_key;unique;not null" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	Password string `json:"password"`
	Address  string `gorm:"type:text" json:"address"`
	Phone    string `json:"phone"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type user []User
