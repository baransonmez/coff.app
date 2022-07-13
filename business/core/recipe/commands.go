package recipe

import (
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
		return &common.CannotBeEmptyError{Field: "user_id"}
	}
	if r.CoffeeID == "" {
		return &common.CannotBeEmptyError{Field: "coffee_id"}
	}
	if r.Description == "" {
		return &common.CannotBeEmptyError{Field: "description"}
	}
	if len(r.Steps) < 1 {
		return &common.CannotBeSmallerError{Field: "steps length", Limit: 1}
	}

	return nil
}
