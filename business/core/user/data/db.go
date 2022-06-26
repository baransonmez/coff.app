package data

import (
	"context"
	"github.com/baransonmez/coff.app/business/core/user"
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
