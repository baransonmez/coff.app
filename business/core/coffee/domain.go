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
