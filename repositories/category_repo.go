package repositories

import (
	db2 "docApp/db"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type CategoryRepo struct {
	Db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{Db: db}
}

func (r *CategoryRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var category []models.Category

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&category)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *CategoryRepo) Create(category *models.Category) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&category)
		return RepositoryResult{Result: category}
	}
}

func (r *CategoryRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var category []models.Category
		db.Find(&category)
		return RepositoryResult{Result: category}
	}
}

func (r *CategoryRepo) Update(category *models.Category) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&category)
		return RepositoryResult{Result: category}
	}
}

func (*CategoryRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var category []models.Category

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&category)
	return RepositoryResult{Result: db.RowsAffected}
}
