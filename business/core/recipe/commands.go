package recipe

import (
	"github.com/baransonmez/coff.app/business/common"
	"time"

	"github.com/google/uuid"
)

type NewRecipe struct {
	UserID      string `json:"user_id" validate:"required"`
	CoffeeID    string `json:"coffee_id" validate:"required"`
	Description string `json:"desc" validate:"required"`
}

func (r NewRecipe) toDomainModel() Recipe {
	userUuidFromString, _ := common.StringToID(r.UserID)
	coffeeUuidFromString, _ := common.StringToID(r.CoffeeID)
	recipe := Recipe{
		ID:          uuid.New(),
		Description: r.Description,
		CoffeeID:    coffeeUuidFromString,
		UserID:      userUuidFromString,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return recipe
}
