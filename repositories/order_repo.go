package repositories

import (
	db2 "docApp/db"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	Db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{Db: db}
}

func (r *OrderRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var orders []models.Order

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&orders)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *OrderRepo) Create(order *models.Order) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&order)
		return RepositoryResult{Result: order}
	}
}

func (r *OrderRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var orders []models.Order
		db.Preload("User").Find(&orders)
		db.Preload("OrderDetail").Find(&orders)
		return RepositoryResult{Result: orders}
	}
}

func (r *OrderRepo) Update(order *models.Order) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&order)
		return RepositoryResult{Result: order}
	}
}

func (*OrderRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var order []models.Order

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&order)
	return RepositoryResult{Result: db.RowsAffected}
}
