package data

import (
	"context"
	"errors"
	"github.com/baransonmez/coff.app/business/core/coffee"
)

type inMem struct {
	m map[coffee.ID]*Bean
}

func NewInMem() *inMem {
	var m = map[coffee.ID]*Bean{}
	return &inMem{
		m: m,
	}
}

func (i *inMem) Create(ctx context.Context, bean coffee.Bean) error {
	coffeeBeanForDB := &Bean{
		ID:          bean.ID.String(),
		Name:        bean.Name,
		Roaster:     bean.Roaster,
		Origin:      bean.Origin,
		RoastDate:   bean.RoastDate,
		DateCreated: bean.DateCreated,
		DateUpdated: bean.DateUpdated,
	}
	i.m[bean.ID] = coffeeBeanForDB
	return nil
}

func (i *inMem) Get(id coffee.ID) (*coffee.Bean, error) {
	if i.m[id] == nil {
		return nil, errors.New("not found")
	}
	return toBean(i.m[id]), nil
}
