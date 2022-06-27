package data

import (
	"context"
	"errors"
	"github.com/baransonmez/coff.app/business/core/user"
	"sync"
)

type inMem struct {
	store map[user.ID]*User
	m     sync.Mutex
}

func NewInMem() *inMem {
	var emptyMap = map[user.ID]*User{}
	return &inMem{
		store: emptyMap,
	}
}

func (i *inMem) Create(_ context.Context, user user.User) error {
	userForDb := &User{
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
	return toUser(i.store[id]), nil
}
