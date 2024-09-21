package dto

import "fmt"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Age < 18 {
		return fmt.Errorf("age is lower than 18")
	}
	return nil
}
