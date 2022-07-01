package recipe

import (
	"errors"
	"github.com/baransonmez/coff.app/business/common"
	"time"

	"github.com/google/uuid"
)

type NewRecipe struct {
	UserID      string `json:"user_id"`
	CoffeeID    string `json:"coffee_id"`
	Description string `json:"desc"`
	Steps       []Step `json:"steps"`
}

var _ common.Command = &NewRecipe{}

func (r NewRecipe) toDomainModel() Recipe {
	userUuidFromString, _ := common.StringToID(r.UserID)
	coffeeUuidFromString, _ := common.StringToID(r.CoffeeID)
	recipe := Recipe{
		ID:          uuid.New(),
		Description: r.Description,
		CoffeeID:    coffeeUuidFromString,
		UserID:      userUuidFromString,
		Steps:       r.Steps,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return recipe
}

func (r *NewRecipe) Validate() error {
	if r.UserID == "" {
		return errors.New("user_id cannot be empty")
	}
	if r.CoffeeID == "" {
		return errors.New("coffee_id cannot be empty")
	}
	if r.Description == "" {
		return errors.New("description cannot be empty")
	}

	if len(r.Steps) < 1 {
		return errors.New("length of steps cannot be smaller than 1")
	}

	return nil
}
