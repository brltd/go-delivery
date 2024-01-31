package domain

import "time"

type Store struct {
	Id          string    `db:"id"`
	Name        string    `db:"name"`
	ImageURL    string    `db:imageURL`
	Icon        string    `db:"icon"`
	StatusId    string    `db:"statusId"`
	DateCreated time.Time `db:"dateCreated"`
	DateUpdated time.Time `db:"dateUpdated"`
}
