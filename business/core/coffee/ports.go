package coffee

import "context"

// Reader interface
type Reader interface {
	Get(id ID) (*Bean, error)
}

// Writer Coffee Bean writer
type Writer interface {
	Create(ctx context.Context, e Bean) (ID error)
	//Update(ctx context.Context, e *Bean) error
	//Delete(ctx context.Context, id int32) error
}

type Repository interface {
	Reader
	Writer
}
