package repositories

import (
	db2 "docApp/db"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type ServiceRepo struct {
	Db *gorm.DB
}

func NewServiceRepo(db *gorm.DB) *ServiceRepo {
	return &ServiceRepo{Db: db}
}

func (r *ServiceRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var services []models.Services

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&services)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *ServiceRepo) Create(service *models.Services) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&service)
		return RepositoryResult{Result: service}
	}
}

func (r *ServiceRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var services []models.Services
		db.Find(&services)
		return RepositoryResult{Result: services}
	}
}

func (r *ServiceRepo) Update(service *models.Services) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&service)
		return RepositoryResult{Result: service}
	}
}

func (*ServiceRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var service []models.Services

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&service)
	return RepositoryResult{Result: db.RowsAffected}
}
