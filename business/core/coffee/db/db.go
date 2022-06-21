package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Store struct {
	log          *zap.SugaredLogger
	db           sqlx.ExtContext
	isWithinTran bool
}

func (s Store) Create(ctx context.Context, bean Bean) error {

	return nil
}
