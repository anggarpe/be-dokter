package repositories

import (
	db2 "docApp/script"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type ProductRepo struct {
	Db *gorm.DB
}
var product []models.Product

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{Db: db}
}

func (r *ProductRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()

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
		//db.Find(&product)
		db.Preload("Category").Find(&product)
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

func (*ProductRepo) Search(name string) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}
	pro := db.Where("name like ?","%"+name+"%").Find(&product)
	if gorm.IsRecordNotFoundError(pro.Error){
		return RepositoryResult{Error: pro.Error}
	}
	return RepositoryResult{Result: pro}
}

func (*ProductRepo) FindAllByPrice(min, max float64) RepositoryResult {
	db, err := db2.DbConn()

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("price >= ? and price <= ?", min, max).Find(&product)
	return RepositoryResult{Result: product}
}

func (*ProductRepo) OrderByPriceDsc() RepositoryResult {
	db, err := db2.DbConn()

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Order("price desc").Find(&product)
	return RepositoryResult{Result: product}
}