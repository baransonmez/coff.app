package persistence

import (
	"context"
	"errors"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"github.com/baransonmez/coff.app/business/core/recipe/data"
	"sync"
)

type inMem struct {
	store map[recipe.ID]*data.Recipe
	m     sync.Mutex
}

func NewInMem() *inMem {
	var emptyMap = map[recipe.ID]*data.Recipe{}
	return &inMem{
		store: emptyMap,
	}
}

func (i *inMem) Create(_ context.Context, recipe recipe.Recipe) error {
	recipeForDb := &data.Recipe{
		ID:          recipe.ID.String(),
		Description: recipe.Description,
		UserID:      recipe.UserID.String(),
		CoffeeID:    recipe.CoffeeID.String(),
		Steps:       data.StepsFromDomainModel(recipe.Steps),
		DateCreated: recipe.DateCreated,
		DateUpdated: recipe.DateUpdated,
	}
	i.m.Lock()
	defer i.m.Unlock()
	i.store[recipe.ID] = recipeForDb
	return nil
}

func (i *inMem) Get(id recipe.ID) (*recipe.Recipe, error) {
	if i.store[id] == nil {
		return nil, errors.New("not found")
	}
	return i.store[id].ToRecipe(), nil
}
