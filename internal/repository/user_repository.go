package repository

import (
	"github.com/mottajunior/glauber-io/internal/schemas"
)

type UserRepository interface {
	Save(user schemas.User) error
}
