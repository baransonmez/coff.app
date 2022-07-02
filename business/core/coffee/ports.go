package coffee

import "context"

type Repository interface {
	Get(id ID) (*Bean, error)
	Create(ctx context.Context, e Bean) (ID error)
}
