package models

import (
	"time"
)

type Reservation struct {
	ID uint `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	ReservationDate *time.Time `json:"reservationDate"`
	Needs string `gorm:"type:text" json:"needs"`
	UserID uint `json:"-"`
	User User `json:"user"`
}
