package data

import (
	"github.com/baransonmez/coff.app/business/core/recipe"
	"github.com/google/uuid"
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
	uuidFromString, _ := StringToID(dbRecipe.ID)
	userUuidFromString, _ := StringToID(dbRecipe.UserID)
	coffeeUuidFromString, _ := StringToID(dbRecipe.CoffeeID)
	recipe := recipe.Recipe{
		ID:          uuidFromString,
		Description: dbRecipe.Description,
		UserID:      userUuidFromString,
		CoffeeID:    coffeeUuidFromString,
		DateCreated: dbRecipe.DateCreated,
		DateUpdated: dbRecipe.DateUpdated,
	}
	return &recipe
}

func StringToID(s string) (recipe.ID, error) {
	id, err := uuid.Parse(s)
	return recipe.ID(id), err
}
