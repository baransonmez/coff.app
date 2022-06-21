package db

import "time"

type Bean struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Roaster     string    `db:"roaster"`
	Origin      string    `db:"origin"`
	Price       int       `db:"price"`
	RoastDate   time.Time `db:"roast_created"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}
