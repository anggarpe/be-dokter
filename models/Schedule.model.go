package models

type Schedule struct {
	ID        uint   `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Day       string `gorm:"not null" json:"day"`
	HourStart string `json:"FromHour"`
	HourEnd   string `json:"ToHour"`
	DoctorID  uint   `sql:"index" json:"idDoctor"`
}
