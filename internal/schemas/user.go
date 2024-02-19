package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}
