package data

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"time"
)

type Recipe struct {
	ID          string    `data:"id"`
	UserID      string    `json:"user_id"`
	CoffeeID    string    `json:"coffee_id"`
	Description string    `json:"desc"`
	DateCreated time.Time `data:"date_created"`
	DateUpdated time.Time `data:"date_updated"`
}

func toRecipe(dbRecipe *Recipe) *recipe.Recipe {
	uuidFromString, _ := common.StringToID(dbRecipe.ID)
	userUuidFromString, _ := common.StringToID(dbRecipe.UserID)
	coffeeUuidFromString, _ := common.StringToID(dbRecipe.CoffeeID)
	dbToDomainModel := recipe.Recipe{
		ID:          uuidFromString,
		Description: dbRecipe.Description,
		UserID:      userUuidFromString,
		CoffeeID:    coffeeUuidFromString,
		DateCreated: dbRecipe.DateCreated,
		DateUpdated: dbRecipe.DateUpdated,
	}
	return &dbToDomainModel
}
