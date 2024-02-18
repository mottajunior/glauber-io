package schemas

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}
