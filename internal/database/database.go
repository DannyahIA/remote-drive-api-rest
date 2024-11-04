package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func ConnectDatabase() error {
	// Define the DSN (Data Source Name)
	dsn := "host=localhost user=postgres password=postgres dbname=remote-drive port=5432 sslmode=disable"

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	Gorm = db

	return nil
}
