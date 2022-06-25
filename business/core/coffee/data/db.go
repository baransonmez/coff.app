package data

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"time"
)

type store struct {
	//log          *zap.SugaredLogger
	//data           sqlx.ExtContext
	//isWithinTran bool
}

func NewStore() store {
	return store{
		//log: log,
		//data:  data,
	}
}

func (s store) Create(ctx context.Context, bean coffee.Bean) error {

	return nil
}

func (s store) Get(id coffee.ID) (*coffee.Bean, error) {
	coffeeBeanDB := &Bean{
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
