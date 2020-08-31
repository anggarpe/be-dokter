package repositories

import (
	"docApp/models"
	db2 "docApp/script"
	"github.com/jinzhu/gorm"
)

type OrderDetailRepo struct {
	Db *gorm.DB
}

func NewOrderDetailRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{Db: db}
}

func (r *OrderDetailRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var detail []models.OrderDetail
		db.Find(&detail)
		return RepositoryResult{Result: detail}
	}
}
