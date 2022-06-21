package coffee

import (
	"github.com/baransonmez/coff.app/business/core/coffee/db"
	"time"
	"unsafe"
)

type Bean struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Roaster     string    `json:"roaster"`
	Origin      string    `json:"origin"`
	Price       int       `json:"price"`
	RoastDate   time.Time `json:"roast_created"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type NewCoffeeBean struct {
	Name      string    `json:"name" validate:"required"`
	Roaster   string    `json:"roaster" validate:"required"`
	Origin    string    `json:"origin" validate:"required"`
	Price     int       `json:"price" validate:"required,gte=0"`
	RoastDate time.Time `json:"roast_created"`
}

func toBean(dbPrd db.Bean) Bean {
	pu := (*Bean)(unsafe.Pointer(&dbPrd))
	return *pu
}
