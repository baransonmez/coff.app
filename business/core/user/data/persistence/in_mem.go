package persistence

import (
	"context"
	"errors"
	"github.com/baransonmez/coff.app/business/core/user"
	"github.com/baransonmez/coff.app/business/core/user/data"
	"sync"
)

type inMem struct {
	store map[user.ID]*data.User
	m     sync.Mutex
}

func NewInMem() *inMem {
	var emptyMap = map[user.ID]*data.User{}
	return &inMem{
		store: emptyMap,
	}
}

func (i *inMem) Create(_ context.Context, user user.User) error {
	userForDb := &data.User{
		ID:          user.ID.String(),
		Name:        user.Name,
		DateCreated: user.DateCreated,
		DateUpdated: user.DateUpdated,
	}
	i.m.Lock()
	defer i.m.Unlock()
	i.store[user.ID] = userForDb
	return nil
}

func (i *inMem) Get(id user.ID) (*user.User, error) {
	if i.store[id] == nil {
		return nil, errors.New("not found")
	}
	return i.store[id].ToUser(), nil
}
