package handler

import "fmt"

type CreateAccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (r *CreateAccountRequest) Validate() error {
	if r.Age < 18 {
		return fmt.Errorf("age is lower than 18")
	}
	return nil
}
