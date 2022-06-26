package user

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

func (c Service) CreateNewUser(ctx context.Context, np NewUser) (ID, error) {
	user := np.toDomainModel()
	if err := c.repository.Create(ctx, user); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return user.ID, nil
}

func (c Service) GetUser(ctx context.Context, id ID) (*User, error) {
	bean, err := c.repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return bean, nil
}
