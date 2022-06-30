package data

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"time"
)

type store struct {
}

func NewStore() store {
	return store{}
}

func (s store) Create(_ context.Context, recipe recipe.Recipe) error {

	return nil
}

func (s store) Get(id recipe.ID) (*recipe.Recipe, error) {
	recipeDB := &Recipe{
		ID:          "uuid.New()",
		Description: "recipe.Description",
		UserID:      "recipe.UserID.String()",
		CoffeeID:    "recipe.CoffeeID.String()",
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}
	return toRecipe(recipeDB), nil
}
