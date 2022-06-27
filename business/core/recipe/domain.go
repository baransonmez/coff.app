package recipe

import (
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/baransonmez/coff.app/business/core/user"
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID
type Recipe struct {
	ID          ID        `json:"id"`
	UserID      user.ID   `json:"user_id"`
	CoffeeID    coffee.ID `json:"coffee_id"`
	Description string    `json:"desc"`
	Steps       []Step    `json:"steps"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type Step struct {
	Description       string
	DurationInSeconds int32
}
