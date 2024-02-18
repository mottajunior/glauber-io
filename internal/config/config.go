package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = ConnectDatabase()
	if err != nil {
		return fmt.Errorf("error when trying to connect to the database: %v", err)
	}

	return nil
}

func GetDatabaseConnection() *gorm.DB {
	return db
}
