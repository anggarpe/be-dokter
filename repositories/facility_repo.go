package repositories

import (
	db2 "docApp/script"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type FacilityRepo struct {
	Db *gorm.DB
}

func NewFacilityRepo(db *gorm.DB) *FacilityRepo {
	return &FacilityRepo{Db: db}
}

var facilities []models.Facility

func (r *FacilityRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&facilities)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *FacilityRepo) Create(Facility *models.Facility) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&Facility)
		return RepositoryResult{Result: Facility}
	}
}

func (r *FacilityRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Find(&facilities)
		return RepositoryResult{Result: facilities}
	}
}

func (r *FacilityRepo) Update(Facility *models.Facility) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&Facility)
		return RepositoryResult{Result: Facility}
	}
}

func (*FacilityRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&facilities)
	return RepositoryResult{Result: db.RowsAffected}
}

func (*FacilityRepo) FindByName(name string) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}
	pro := db.Where("name like ?","%"+name+"%").Find(&facilities)
	if gorm.IsRecordNotFoundError(pro.Error){
		return RepositoryResult{Error: pro.Error}
	}
	return RepositoryResult{Result: &facilities}
}
