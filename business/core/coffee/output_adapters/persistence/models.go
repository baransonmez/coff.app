package persistence

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"time"
)

type Bean struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Roaster     string    `db:"roaster"`
	Origin      string    `db:"origin"`
	Price       int       `db:"price"`
	RoastDate   time.Time `db:"roast_date"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
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
