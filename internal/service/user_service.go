package service

import (
	"myapp/internal/model"
	"myapp/internal/repository"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB, limit int) ([]model.User, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	return repository.FindAllUsers(db, limit)
}

func RegisterUser(db *gorm.DB, user *model.User) error {
	return repository.CreateUser(db, user)
}
