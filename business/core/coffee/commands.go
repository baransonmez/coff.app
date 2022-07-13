package coffee

import (
	"github.com/baransonmez/coff.app/business/common"
	"time"

	"github.com/google/uuid"
)

type NewCoffeeBean struct {
	Name      string    `json:"name"`
	Roaster   string    `json:"roaster"`
	Origin    string    `json:"origin"`
	Price     int       `json:"price"`
	RoastDate time.Time `json:"roast_created"`
}

var _ common.Command = &NewCoffeeBean{}

func (c NewCoffeeBean) toDomainModel() *Bean {
	coffeeBean := Bean{
		ID:          uuid.New(),
		Name:        c.Name,
		Roaster:     c.Roaster,
		Origin:      c.Origin,
		RoastDate:   c.RoastDate,
		Price:       c.Price,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return &coffeeBean
}

func (c *NewCoffeeBean) Validate() error {
	if c.Name == "" {
		return &common.CannotBeEmptyError{Field: "Name"}
	}
	if c.Origin == "" {
		return &common.CannotBeEmptyError{Field: "Origin"}
	}
	if c.Roaster == "" {
		return &common.CannotBeEmptyError{Field: "Roaster"}
	}
	if c.Price < 1 {
		return &common.CannotBeSmallerError{Field: "Price", Limit: 1}
	}

	return nil
}
