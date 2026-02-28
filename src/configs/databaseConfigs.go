package configs

import (
	"fmt"

	"segmenta/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDatabase(currentConfig *models.AppConfig) error {
	dsn := currentConfig.DBConnection
	if dsn == "" {
		sslMode := currentConfig.DBSSLMode
		if sslMode == "" {
			sslMode = "disable"
		}
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			currentConfig.DBHost,
			currentConfig.DBPort,
			currentConfig.DBUser,
			currentConfig.DBPassword,
			currentConfig.DBName,
			sslMode,
		)
	}

	db, errorHandler := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorHandler != nil {
		return fmt.Errorf("[DB] Failed to connect to database: %w", errorHandler)
	}

	errorHandler = db.AutoMigrate(&models.User{})
	if errorHandler != nil {
		return fmt.Errorf("[DB] Failed to auto migrate: %w", errorHandler)
	}

	Database = db
	return nil
}
