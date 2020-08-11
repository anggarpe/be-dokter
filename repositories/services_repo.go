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

var services []models.Services

func (r *ServiceRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()

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
	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&services)
	return RepositoryResult{Result: db.RowsAffected}
}

func (*ServiceRepo) FindByName(name string) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}
	pro := db.Where("name like ?","%"+name+"%").Find(&services)
	if gorm.IsRecordNotFoundError(pro.Error){
		return RepositoryResult{Error: pro.Error}
	}
	return RepositoryResult{Result: &services}
}
