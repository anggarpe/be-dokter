package db

import (
	"docApp/models"
	"github.com/jinzhu/gorm"
)

func CreateConstraint(db *gorm.DB)  {
	db.Model(&models.Schedule{}).AddForeignKey("id_doctor", "doctors(id)", "cascade", "cascade")
}
