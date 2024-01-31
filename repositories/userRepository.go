package repositories

import (
	"github.com/brltd/delivery/domain"
	"github.com/brltd/delivery/helpers"
	"github.com/brltd/delivery/logger"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(user domain.User) (*domain.User, *helpers.ApiError)
}

type UserRepositoryDb struct {
	client *sqlx.DB
}

func NewUserRepository(dbClient *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{client: dbClient}
}

func (u UserRepositoryDb) CreateUser(user domain.User) (*domain.User, *helpers.ApiError) {
	stmt := `
		INSERT INTO users ("name", email, hash, "dateCreated", "dateUpdated")
			VALUES($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id string

	err := u.client.Get(&id, stmt,
		user.Name,
		user.Email,
		user.Hash,
		user.DateCreated,
		user.DateUpdated)

	if err != nil {
		pqErr, isPqError := err.(*pq.Error)

		// Duplicate key violation
		if isPqError && pqErr.Code == "23505" {

			return nil, helpers.ValidationError("E-mail already exists")
		}

		logger.Error("Error while creating new user " + err.Error())
		return nil, helpers.InternalServerError("Unexpected error from database")
	}

	user.Id = id

	return &user, nil
}
