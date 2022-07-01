package user

import (
	"errors"
	"github.com/baransonmez/coff.app/business/common"
	"time"

	"github.com/google/uuid"
)

type NewUser struct {
	Name string `json:"name"`
}

var _ common.Command = &NewUser{}

func (u NewUser) toDomainModel() User {
	user := User{
		ID:          uuid.New(),
		Name:        u.Name,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return user
}

func (u *NewUser) Validate() error {
	if u.Name == "" {
		return errors.New("user name cannot be empty")
	}
	return nil
}
