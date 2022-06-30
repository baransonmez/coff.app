package coffee

import (
	"errors"
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

func (b *Bean) validate() error {
	if b.Name == "" {
		return errors.New("bean name cannot be empty")
	}
	if b.Origin == "" {
		return errors.New("bean origin cannot be empty")
	}
	if b.Roaster == "" {
		return errors.New("bean roaster cannot be empty")
	}

	if b.Price < 1 {
		return errors.New("bean price cannot be smaller than 1")
	}

	return nil
}
