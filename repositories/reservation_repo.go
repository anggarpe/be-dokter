package repositories

import (
	db2 "docApp/script"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type ReservationRepo struct {
	Db *gorm.DB
}

func NewReservationRepo(db *gorm.DB) *ReservationRepo {
	return &ReservationRepo{Db: db}
}

func (r *ReservationRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var reservations []models.Reservation

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&reservations)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *ReservationRepo) Create(reservation *models.Reservation) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&reservation)
		return RepositoryResult{Result: reservation}
	}
}

func (r *ReservationRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var reservations []models.Reservation
		db.Preload("User").Preload("Variety").Find(&reservations)
		return RepositoryResult{Result: reservations}
	}
}

func (r *ReservationRepo) Update(reservation *models.Reservation) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&reservation)
		return RepositoryResult{Result: reservation}
	}
}

func (*ReservationRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var reservation []models.Reservation

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&reservation)
	return RepositoryResult{Result: db.RowsAffected}
}
