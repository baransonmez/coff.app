package persistence

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/user"
	"time"
)

type User struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
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
