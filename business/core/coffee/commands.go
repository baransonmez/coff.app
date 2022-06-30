package coffee

import (
	"time"

	"github.com/google/uuid"
)

type NewCoffeeBean struct {
	Name      string    `json:"name" validate:"required"`
	Roaster   string    `json:"roaster" validate:"required"`
	Origin    string    `json:"origin" validate:"required"`
	Price     int       `json:"price" validate:"required,gte=0"`
	RoastDate time.Time `json:"roast_created"`
}

func (c NewCoffeeBean) toDomainModel() (*Bean, error) {
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
	err := coffeeBean.validate()
	if err != nil {
		return nil, err
	}
	return &coffeeBean, nil
}
