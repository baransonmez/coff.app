package common

import "fmt"

type Command interface {
	Validate() error
}

type CannotBeEmptyError struct {
	Field string
}

func (e *CannotBeEmptyError) Error() string {
	return fmt.Sprintf("%q cannot be empty", e.Field)
}

type CannotBeSmallerError struct {
	Field string
	Limit int
}

func (e *CannotBeSmallerError) Error() string {
	return fmt.Sprintf("%q cannot be smaller than %d", e.Field, e.Limit)
}
