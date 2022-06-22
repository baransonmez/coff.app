package db

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"time"
)

type Store struct {
	//log          *zap.SugaredLogger
	//db           sqlx.ExtContext
	//isWithinTran bool
}

func NewStore() Store {
	return Store{
		//log: log,
		//db:  db,
	}
}

func (s Store) Create(ctx context.Context, bean coffee.Bean) error {

	return nil
}

func (s Store) Get(id coffee.ID) (coffee.Bean, error) {
	coffeeBeanDB := Bean{
		ID:          "uuid.New()",
		Name:        "np.Name",
		Roaster:     "np.Roaster",
		Origin:      "np.Origin",
		RoastDate:   time.Now(),
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}
	return toBean(coffeeBeanDB), nil
}
