package db

import (
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/google/uuid"
	"time"
)

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

func toBean(dbPrd Bean) coffee.Bean {
	uuidFromString, _ := StringToID(dbPrd.ID)
	pu := coffee.Bean{
		ID:          uuidFromString,
		Name:        dbPrd.Name,
		Roaster:     dbPrd.Roaster,
		Origin:      dbPrd.Origin,
		RoastDate:   dbPrd.RoastDate,
		DateCreated: dbPrd.DateCreated,
		DateUpdated: dbPrd.DateUpdated,
	}
	return pu
}

func StringToID(s string) (coffee.ID, error) {
	id, err := uuid.Parse(s)
	return coffee.ID(id), err
}
