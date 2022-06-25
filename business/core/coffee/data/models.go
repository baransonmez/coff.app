package data

import (
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/google/uuid"
	"time"
)

type Bean struct {
	ID          string    `data:"id"`
	Name        string    `data:"name"`
	Roaster     string    `data:"roaster"`
	Origin      string    `data:"origin"`
	Price       int       `data:"price"`
	RoastDate   time.Time `data:"roast_created"`
	DateCreated time.Time `data:"date_created"`
	DateUpdated time.Time `data:"date_updated"`
}

func toBean(dbPrd *Bean) *coffee.Bean {
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
	return &pu
}

func StringToID(s string) (coffee.ID, error) {
	id, err := uuid.Parse(s)
	return coffee.ID(id), err
}
