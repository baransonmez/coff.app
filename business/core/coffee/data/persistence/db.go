package persistence

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/baransonmez/coff.app/business/core/coffee/data"
	"time"
)

type store struct {
}

func NewStore() store {
	return store{}
}

func (s store) Create(_ context.Context, bean coffee.Bean) error {

	return nil
}

func (s store) Get(id coffee.ID) (*coffee.Bean, error) {
	coffeeBeanDB := &data.Bean{
		ID:          "uuid.New()",
		Name:        "np.Name",
		Roaster:     "np.Roaster",
		Origin:      "np.Origin",
		RoastDate:   time.Now(),
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}
	return coffeeBeanDB.ToBean(), nil
}
