package repositories

import (
	db2 "docApp/db"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type DoctorRepo struct {
	Db *gorm.DB
}

func NewDoctorRepo(db *gorm.DB) *DoctorRepo {
	return &DoctorRepo{Db: db}
}

func (r *DoctorRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var doctors []models.Doctor

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&doctors)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *DoctorRepo) Create(doctor *models.Doctor) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&doctor)
		return RepositoryResult{Result: doctor}
	}
}

func (r *DoctorRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var doctors []models.Doctor
		db.Preload("Schedule").Find(&doctors)
		//db.Find(&doctors)
		return RepositoryResult{Result: doctors}
	}
}

func (r *DoctorRepo) Update(doctor *models.Doctor) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&doctor)
		return RepositoryResult{Result: doctor}
	}
}

func (*DoctorRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var doctor []models.Doctor

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&doctor)
	return RepositoryResult{Result: db.RowsAffected}
}
