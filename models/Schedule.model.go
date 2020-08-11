package models

type Schedule struct {
	ID        uint   `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Day       string `gorm:"not null" json:"day"`
	StartHour string `json:"from_hour"`
	EndHour   string `json:"to_hour"`
	DoctorID  uint `json:"-"`
	Doctor Doctor `gorm:"association_autoupdate:false" json:"-"`
}
