package recipe

import "context"

// Reader interface
type Reader interface {
	Get(id ID) (*Recipe, error)
}

// Writer Recipe writer
type Writer interface {
	Create(ctx context.Context, e Recipe) (ID error)
}

type Repository interface {
	Reader
	Writer
}
