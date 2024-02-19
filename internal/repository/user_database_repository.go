package repository

import (
	"log/slog"

	"github.com/mottajunior/glauber-io/internal/schemas"
	"gorm.io/gorm"
)

type UserDatabaseRepository struct {
	db *gorm.DB
}

func NewUserDatabaseRepository(db *gorm.DB) UserRepository {
	repository := UserDatabaseRepository{db: db}
	return repository
}

func (a UserDatabaseRepository) Save(user schemas.User) error {
	if err := a.db.Create(&user).Error; err != nil {
		slog.Error("Error when save user in database: ", err)
		return err
	}
	return nil
}
