package user

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID
type User struct {
	ID          ID        `json:"id"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
