package repository

import (
	"myapp/internal/model"

	"gorm.io/gorm"
)

func FindAllUsers(db *gorm.DB, limit int) ([]model.User, error) {
	var users []model.User
	if err := db.Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}
