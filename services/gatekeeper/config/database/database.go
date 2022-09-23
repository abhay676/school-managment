package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URI")), &gorm.Config{})

	if err != nil {
		fmt.Println("[DATABASE]:: Connection Error")
		panic(err)
	}

	DB = db
	fmt.Println("[DATABASE]::Connected")
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
