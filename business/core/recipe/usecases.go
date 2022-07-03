package recipe

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Service struct {
	recipeRepository Repository
	userRepository   UserRepository
}

func NewService(r Repository, u UserRepository) *Service {
	return &Service{
		recipeRepository: r,
		userRepository:   u,
	}
}

func (c Service) CreateNewRecipe(ctx context.Context, np NewRecipe) (ID, error) {
	err := np.Validate()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	_, err = c.userRepository.IsValidUser(ctx, np.UserID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	recipe := np.toDomainModel()
	if err := c.recipeRepository.Create(ctx, recipe); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return recipe.ID, nil
}

func (c Service) GetRecipe(_ context.Context, id ID) (*Recipe, error) {
	recipe, err := c.recipeRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return recipe, nil
}
