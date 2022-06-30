package coffee

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

func (c Service) CreateCoffeeBean(ctx context.Context, np NewCoffeeBean) (ID, error) {
	coffeeBean, err := np.toDomainModel()
	if err != nil {
		return uuid.UUID{}, err
	}
	if err := c.repository.Create(ctx, *coffeeBean); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return coffeeBean.ID, nil
}

func (c Service) GetCoffeeBean(_ context.Context, id ID) (*Bean, error) {
	bean, err := c.repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return bean, nil
}
