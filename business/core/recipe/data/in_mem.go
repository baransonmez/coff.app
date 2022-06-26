package data

import (
	"context"
	"errors"
	"github.com/baransonmez/coff.app/business/core/recipe"
)

type inMem struct {
	m map[recipe.ID]*Recipe
}

func NewInMem() *inMem {
	var m = map[recipe.ID]*Recipe{}
	return &inMem{
		m: m,
	}
}

func (i *inMem) Create(_ context.Context, recipe recipe.Recipe) error {
	recipeForDb := &Recipe{
		ID:          recipe.ID.String(),
		Description: recipe.Description,
		UserID:      recipe.UserID.String(),
		CoffeeID:    recipe.CoffeeID.String(),
		DateCreated: recipe.DateCreated,
		DateUpdated: recipe.DateUpdated,
	}
	i.m[recipe.ID] = recipeForDb
	return nil
}

func (i *inMem) Get(id recipe.ID) (*recipe.Recipe, error) {
	if i.m[id] == nil {
		return nil, errors.New("not found")
	}
	return toRecipe(i.m[id]), nil
}
