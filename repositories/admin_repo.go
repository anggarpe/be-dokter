package repositories

import (
	db2 "docApp/script"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type AdminRepo struct {
	Db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{Db: db}
}

func (r *AdminRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var admins []models.Admin

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&admins)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *AdminRepo) Create(admin *models.Admin) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&admin)
		return RepositoryResult{Result: admin}
	}
}

func (r *AdminRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var admins []models.Admin
		db.Find(&admins)
		return RepositoryResult{Result: admins}
	}
}

func (r *AdminRepo) Update(admin *models.Admin) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&admin)
		return RepositoryResult{Result: admin}
	}
}

func (*AdminRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var admin []models.Admin

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&admin)
	return RepositoryResult{Result: db.RowsAffected}
}
