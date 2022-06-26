package recipe

import "context"

// Reader interface
type Reader interface {
	Get(id ID) (*Recipe, error)
}

// Writer Recipe writer
type Writer interface {
	Create(ctx context.Context, e Recipe) (ID error)
	//Update(ctx context.Context, e *Recipe) error
	//Delete(ctx context.Context, id int32) error
}

type Repository interface {
	Reader
	Writer
}
