package repositories

import (
	"github.com/jinzhu/gorm"
	"docApp/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Save(user *models.User) RepositoryResult {
	err := r.db.Save(user).Error

	if err != nil{
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: user}
}