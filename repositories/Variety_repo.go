package repositories

import (
	db2 "docApp/script"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type VarietyRepo struct {
	Db *gorm.DB
}

func NewVarietyRepo(db *gorm.DB) *VarietyRepo {
	return &VarietyRepo{Db: db}
}

func (r *VarietyRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var varietys []models.Variety

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&varietys)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *VarietyRepo) Create(variety *models.Variety) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&variety)
		return RepositoryResult{Result: variety}
	}
}

func (r *VarietyRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var varietys []models.Variety
		db.Find(&varietys)
		return RepositoryResult{Result: varietys}
	}
}

func (r *VarietyRepo) Update(variety *models.Variety) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&variety)
		return RepositoryResult{Result: variety}
	}
}

func (*VarietyRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var variety []models.Variety

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&variety)
	return RepositoryResult{Result: db.RowsAffected}
}
