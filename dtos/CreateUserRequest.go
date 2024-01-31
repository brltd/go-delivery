package dtos

import "github.com/brltd/delivery/helpers"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c CreateUserRequest) Validate() *helpers.ApiError {
	if c.Name == "" {
		return helpers.ValidationError("Name must not be empty")
	}

	if c.Email == "" {
		return helpers.ValidationError("Email must not be empty")
	}

	if c.Password == "" {
		return helpers.ValidationError("Password must not be empty")
	}

	return nil
}
