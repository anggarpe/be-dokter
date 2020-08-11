package models

type Variety struct {
	ID uint `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Reservation []Reservation `gorm:"many2many:reservation_variety"`
}
