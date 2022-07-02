package user

import "context"

type Repository interface {
	Get(id ID) (*User, error)
	Create(ctx context.Context, e User) (ID error)
}
