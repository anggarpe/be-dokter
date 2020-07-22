package repositories

import (
	db2 "docApp/db"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type ProductRepo struct {
	Db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{Db: db}
}

func (r *ProductRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var product []models.Product

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&product)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *ProductRepo) Create(product *models.Product) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&product)
		return RepositoryResult{Result: product}
	}
}

func (r *ProductRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var product []models.Product
		db.Find(&product)
		return RepositoryResult{Result: product}
	}
}

func (r *ProductRepo) Update(product *models.Product) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&product)
		return RepositoryResult{Result: product}
	}
}

func (*ProductRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var product []models.Product

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&product)
	return RepositoryResult{Result: db.RowsAffected}
}
