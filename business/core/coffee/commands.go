package coffee

import (
	"errors"
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
		return errors.New("bean name cannot be empty")
	}
	if c.Origin == "" {
		return errors.New("bean origin cannot be empty")
	}
	if c.Roaster == "" {
		return errors.New("bean roaster cannot be empty")
	}

	if c.Price < 1 {
		return errors.New("bean price cannot be smaller than 1")
	}

	return nil
}
