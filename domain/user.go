package domain

import (
	"time"

	"github.com/brltd/delivery/dtos"
)

type User struct {
	Id          string     `db:"id"`
	Name        string     `db:"name"`
	Email       string     `db:"email"`
	Hash        string     `db:"hash"`
	DateCreated time.Time  `db:"dateCreated"`
	DateUpdated *time.Time `db:"DateUpdated"`
}

func (u User) ToCreateUserResponseDto() dtos.CreateUserResponse {
	return dtos.CreateUserResponse{
		Id: u.Id,
	}
}
