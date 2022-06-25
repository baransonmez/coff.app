package data

import (
	"github.com/baransonmez/coff.app/business/core/user"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          string    `data:"id"`
	Name        string    `data:"name"`
	DateCreated time.Time `data:"date_created"`
	DateUpdated time.Time `data:"date_updated"`
}

func toUser(dbPrd *User) *user.User {
	uuidFromString, _ := StringToID(dbPrd.ID)
	pu := user.User{
		ID:          uuidFromString,
		Name:        dbPrd.Name,
		DateCreated: dbPrd.DateCreated,
		DateUpdated: dbPrd.DateUpdated,
	}
	return &pu
}

func StringToID(s string) (user.ID, error) {
	id, err := uuid.Parse(s)
	return user.ID(id), err
}
