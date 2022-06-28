package data

import (
	"context"
	"errors"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"sync"
)

type inMem struct {
	store map[coffee.ID]*Bean
	m     sync.Mutex
}

func NewInMem() *inMem {
	var emptyMap = map[coffee.ID]*Bean{}
	return &inMem{
		store: emptyMap,
	}
}

func (i *inMem) Create(_ context.Context, bean coffee.Bean) error {
	coffeeBeanForDB := &Bean{
		ID:          bean.ID.String(),
		Name:        bean.Name,
		Roaster:     bean.Roaster,
		Origin:      bean.Origin,
		Price:       bean.Price,
		RoastDate:   bean.RoastDate,
		DateCreated: bean.DateCreated,
		DateUpdated: bean.DateUpdated,
	}
	i.m.Lock()
	defer i.m.Unlock()
	i.store[bean.ID] = coffeeBeanForDB
	return nil
}

func (i *inMem) Get(id coffee.ID) (*coffee.Bean, error) {
	if i.store[id] == nil {
		return nil, errors.New("not found")
	}
	return toBean(i.store[id]), nil
}
