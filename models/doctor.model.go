package models

type Doctor struct {
	ID       uint  `gorm:"type:int(5);primary_key;unique;not null" json:"id"`
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	Phone    string `gorm:"type:varchar(12)" json:"phone"`
	Schedule []Schedule  `gorm:"foreign_key:DoctorID" json:"schedule"`
}
