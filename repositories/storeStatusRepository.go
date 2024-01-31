package repositories

import "github.com/jmoiron/sqlx"

type StoreStatusRepository interface {
}

type StoreStatusRepositoryDb struct {
	client *sqlx.DB
}
