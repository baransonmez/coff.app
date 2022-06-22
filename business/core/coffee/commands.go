package coffee

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (c Service) CreateCoffeeBean(ctx context.Context, np NewCoffeeBean) (ID, error) {
	coffeeBean := Bean{
		ID:          uuid.New(),
		Name:        np.Name,
		Roaster:     np.Roaster,
		Origin:      np.Origin,
		RoastDate:   np.RoastDate,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	if err := c.repository.Create(ctx, &coffeeBean); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return coffeeBean.ID, nil
}
