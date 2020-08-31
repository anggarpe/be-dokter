package repositories

import (
	db2 "docApp/script"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type ScheduleRepo struct {
	Db *gorm.DB
}

func NewScheduleRepo(db *gorm.DB) *ScheduleRepo {
	return &ScheduleRepo{Db: db}
}

var schedules []models.Schedule
func (r *ScheduleRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&schedules)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *ScheduleRepo) Create(schedule *models.Schedule) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&schedule)
		return RepositoryResult{Result: schedule}
	}
}

func (r *ScheduleRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}
		//db.Model(&schedules).Association("DoctorID").Find(&models.Doctor{})
		db.Preload("Doctor").Find(&schedules)
		return RepositoryResult{Result: schedules}

}

func (r *ScheduleRepo) Update(schedule *models.Schedule) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&schedule)
		return RepositoryResult{Result: schedule}
	}
}

func (*ScheduleRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var schedule []models.Schedule

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&schedule)
	return RepositoryResult{Result: db.RowsAffected}
}
