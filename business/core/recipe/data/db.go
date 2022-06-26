package data

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"time"
)

type store struct {
	//log          *zap.SugaredLogger
	//data           sqlx.ExtContext
	//isWithinTran bool
}

func NewStore() store {
	return store{
		//log: log,
		//data:  data,
	}
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
