package coffee

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID
type Bean struct {
	ID          ID        `json:"id"`
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
