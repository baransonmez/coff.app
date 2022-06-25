package user

import (
	"time"

	"github.com/google/uuid"
)

type NewUser struct {
	Name string `json:"name" validate:"required"`
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
