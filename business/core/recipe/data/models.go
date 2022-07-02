package data

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"time"
)

type Recipe struct {
	ID          string    `data:"id"`
	UserID      string    `data:"user_id"`
	CoffeeID    string    `data:"coffee_id"`
	Description string    `data:"desc"`
	Steps       []Step    `data:"steps"`
	DateCreated time.Time `data:"date_created"`
	DateUpdated time.Time `data:"date_updated"`
}

type Step struct {
	Description       string `data:"desc"`
	DurationInSeconds int32  `data:"duration"`
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
