package persistence

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"github.com/baransonmez/coff.app/business/core/recipe/data"
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
	recipeDB := &data.Recipe{
		ID:          "uuid.New()",
		Description: "recipe.Description",
		UserID:      "recipe.UserID.String()",
		CoffeeID:    "recipe.CoffeeID.String()",
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}
	return recipeDB.ToRecipe(), nil
}
