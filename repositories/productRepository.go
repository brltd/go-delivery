package repositories

import "github.com/jmoiron/sqlx"

type ProductRepository interface {
}

type ProductRepositoryDb struct {
	client *sqlx.DB
}
