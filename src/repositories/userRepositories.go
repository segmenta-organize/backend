package repositories

import (
	"segmenta/src/configs"
	"segmenta/src/models"
)

func CreateUser(user *models.User) error {
	return configs.Database.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if errorHandler := configs.Database.Where("email = ?", email).First(&user).Error; errorHandler != nil {
		return nil, errorHandler
	}
	return &user, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if errorHandler := configs.Database.First(&user, id).Error; errorHandler != nil {
		return nil, errorHandler
	}
	return &user, nil
}