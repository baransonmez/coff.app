package data

import (
	"context"
	"errors"
	"github.com/baransonmez/coff.app/business/core/user"
)

type inMem struct {
	m map[user.ID]*User
}

func NewInMem() *inMem {
	var m = map[user.ID]*User{}
	return &inMem{
		m: m,
	}
}

func (i *inMem) Create(ctx context.Context, user user.User) error {
	userForDb := &User{
		ID:          user.ID.String(),
		Name:        user.Name,
		DateCreated: user.DateCreated,
		DateUpdated: user.DateUpdated,
	}
	i.m[user.ID] = userForDb
	return nil
}

func (i *inMem) Get(id user.ID) (*user.User, error) {
	if i.m[id] == nil {
		return nil, errors.New("not found")
	}
	return toUser(i.m[id]), nil
}
