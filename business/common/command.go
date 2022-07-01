package common

type Command interface {
	Validate() error
}
