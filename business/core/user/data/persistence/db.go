package persistence

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/user"
	"github.com/baransonmez/coff.app/business/core/user/data"
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
	userDB := &data.User{
		ID:          "uuid.New()",
		Name:        "np.Name",
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}
	return userDB.ToUser(), nil
}
