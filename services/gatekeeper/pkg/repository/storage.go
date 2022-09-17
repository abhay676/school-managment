package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBConnection(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// db.AutoMigrate() // Import Models
	return db, nil
}
