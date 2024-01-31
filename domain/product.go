package domain

import "time"

type Product struct {
	Id             string     `db:"id"`
	Name           string     `db:"name"`
	ImageURL       string     `db:"imageURL"`
	Price          float64    `db:"price"`
	PromotionPrice *float64   `db:"promotionPrice"`
	CategoryId     string     `db:"categoryId"`
	StoreId        string     `db:"storeId"`
	DateCreated    time.Time  `db:"dateCreated"`
	DateUpdated    *time.Time `db:"dateUpdated"`
}
