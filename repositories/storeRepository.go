package repositories

import "github.com/jmoiron/sqlx"

type StoreRepository interface {
}

type StoreRepositoryDb struct {
	client *sqlx.DB
}
