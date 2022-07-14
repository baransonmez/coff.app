package persistence

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"time"
)

type Recipe struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	CoffeeID    string    `db:"coffee_id"`
	Description string    `db:"description"`
	Steps       []Step    `db:"step"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

type Step struct {
	RecipeID          string    `db:"recipe_id"`
	StepOrder         string    `db:"step_order"`
	Description       string    `db:"description"`
	DurationInSeconds int32     `db:"duration_in_seconds"`
	DateCreated       time.Time `db:"date_created"`
	DateUpdated       time.Time `db:"date_updated"`
}

func (dbRecipe *Recipe) ToRecipe() *recipe.Recipe {
	uuidFromString, _ := common.StringToID(dbRecipe.ID)
	userUuidFromString, _ := common.StringToID(dbRecipe.UserID)
	coffeeUuidFromString, _ := common.StringToID(dbRecipe.CoffeeID)
	dbToDomainModel := recipe.Recipe{
		ID:          uuidFromString,
		Description: dbRecipe.Description,
		UserID:      userUuidFromString,
		CoffeeID:    coffeeUuidFromString,
		Steps:       stepsToDomainModel(dbRecipe.Steps),
		DateCreated: dbRecipe.DateCreated,
		DateUpdated: dbRecipe.DateUpdated,
	}
	return &dbToDomainModel
}

func stepsToDomainModel(steps []Step) []recipe.Step {
	var stepsVO []recipe.Step
	for _, s := range steps {
		stepsVO = append(stepsVO, s.stepToDomainModel())
	}
	return stepsVO
}

func (s Step) stepToDomainModel() recipe.Step {
	return recipe.Step{
		Description:       s.Description,
		DurationInSeconds: s.DurationInSeconds,
		Order:             s.StepOrder,
	}
}

func StepsFromDomainModel(stepsDO []recipe.Step) []Step {
	var stepsDB []Step
	for _, s := range stepsDO {
		stepsDB = append(stepsDB, stepFromDomainModel(s))
	}
	return stepsDB
}

func stepFromDomainModel(s recipe.Step) Step {
	return Step{
		Description:       s.Description,
		DurationInSeconds: s.DurationInSeconds,
	}
}
