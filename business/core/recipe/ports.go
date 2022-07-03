package recipe

import (
	"context"
)

type Repository interface {
	Get(id ID) (*Recipe, error)
	Create(ctx context.Context, e Recipe) (ID error)
}

type UserRepository interface {
	IsValidUser(ctx context.Context, userId string) (bool, error)
}
