package domain

import "time"

type Category struct {
	Id          string     `db:"id"`
	Name        string     `db:"name"`
	StoreId     string     `db:"storeId"`
	DateCreated time.Time  `db:"dateCreated"`
	DateUpdated *time.Time `db:"dateUpdated"`
}
