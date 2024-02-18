package handler

import (
	"github.com/mottajunior/glauber-io/internal/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	db = config.GetDatabaseConnection()
}
