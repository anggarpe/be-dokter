package repositories

import (
	db2 "docApp/script"
	"docApp/models"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (r *UserRepo) FindById(id string) RepositoryResult {
	db, err := db2.DbConn()
	var users []models.User

	if err != nil{
		return RepositoryResult{Error: err}
	}

	usr := db.Where("id=?",id).Find(&users)
	if gorm.IsRecordNotFoundError(usr.Error){
		return RepositoryResult{Error: usr.Error}
	}
	return RepositoryResult{Result: usr}
}

func (r *UserRepo) Create(user *models.User) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Create(&user)
		return RepositoryResult{Result: user}
	}
}

func (r *UserRepo) FindAll() RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		var users []models.User
		db.Find(&users)
		return RepositoryResult{Result: users}
	}
}

func (r *UserRepo) Update(user *models.User) RepositoryResult {
	db, err := db2.DbConn()
	if err != nil {
		return RepositoryResult{Error: err}
	}else {
		db.Save(&user)
		return RepositoryResult{Result: user}
	}
}

func (*UserRepo) Delete(id string) RepositoryResult {
	db, err := db2.DbConn()
	var user []models.User

	if err != nil {
		return RepositoryResult{Error: err}
	}

	db.Where("id = ?", id).Delete(&user)
	return RepositoryResult{Result: db.RowsAffected}
}
