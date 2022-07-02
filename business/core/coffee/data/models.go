package data

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/coffee"
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

func (dbPrd *Bean) ToBean() *coffee.Bean {
	uuidFromString, _ := common.StringToID(dbPrd.ID)
	pu := coffee.Bean{
		ID:          uuidFromString,
		Name:        dbPrd.Name,
		Roaster:     dbPrd.Roaster,
		Origin:      dbPrd.Origin,
		RoastDate:   dbPrd.RoastDate,
		Price:       dbPrd.Price,
		DateCreated: dbPrd.DateCreated,
		DateUpdated: dbPrd.DateUpdated,
	}
	return &pu
}
