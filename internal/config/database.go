package config

import (
	"github.com/mottajunior/glauber-io/internal/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Account{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
