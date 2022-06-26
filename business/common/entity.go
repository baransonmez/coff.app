package common

import (
	"github.com/google/uuid"
)

func StringToID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	return id, err
}
