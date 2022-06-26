package recipe

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (c Service) CreateNewRecipe(ctx context.Context, np NewRecipe) (ID, error) {
	recipe := np.toDomainModel()
	if err := c.repository.Create(ctx, recipe); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return recipe.ID, nil
}

func (c Service) GetRecipe(_ context.Context, id ID) (*Recipe, error) {
	recipe, err := c.repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return recipe, nil
}
