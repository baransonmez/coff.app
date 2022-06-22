package coffee

import "context"

// Reader interface
type Reader interface {
	Get(id int32) (Bean, error)
}

// Writer Coffee Bean writer
type Writer interface {
	Create(ctx context.Context, e *Bean) (int32 error)
	//Update(ctx context.Context, e *Bean) error
	//Delete(ctx context.Context, id int32) error
}

type Repository interface {
	Reader
	Writer
}
