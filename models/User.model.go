package models

import "github.com/labstack/echo"

type User struct {
	ID       int    `gorm:"primary_key;unique;not null" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `json:"password"`
	Address  string `gorm:"index:addr" json:"address"`
	Phone    string `json:"phone"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser(c echo.Context) error {
}
