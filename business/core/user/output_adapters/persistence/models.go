package persistence

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/user"
	"time"
)

type User struct {
	ID          string    `data:"id"`
	Name        string    `data:"name"`
	DateCreated time.Time `data:"date_created"`
	DateUpdated time.Time `data:"date_updated"`
}

func (dbPrd *User) ToUser() *user.User {
	uuidFromString, _ := common.StringToID(dbPrd.ID)
	pu := user.User{
		ID:          uuidFromString,
		Name:        dbPrd.Name,
		DateCreated: dbPrd.DateCreated,
		DateUpdated: dbPrd.DateUpdated,
	}
	return &pu
}
