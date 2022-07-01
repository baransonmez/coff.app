package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type NewUser struct {
	Name string `json:"name"`
}

func (u NewUser) toDomainModel() User {
	user := User{
		ID:          uuid.New(),
		Name:        u.Name,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return user
}

func (u *NewUser) validate() error {
	if u.Name == "" {
		return errors.New("user name cannot be empty")
	}
	return nil
}
