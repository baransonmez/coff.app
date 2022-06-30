package data

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/user"
	"time"
)

type store struct {
}

func NewStore() store {
	return store{}
}

func (s store) Create(_ context.Context, user user.User) error {

	return nil
}

func (s store) Get(id user.ID) (*user.User, error) {
	userDB := &User{
		ID:          "uuid.New()",
		Name:        "np.Name",
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}
	return toUser(userDB), nil
}
