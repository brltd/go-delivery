package repositories

import "github.com/jmoiron/sqlx"

type CategoryRepository interface {
}

type CategoryRepositoryDb struct {
	client *sqlx.DB
}
