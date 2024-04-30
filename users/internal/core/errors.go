package core

import "errors"

var (
	ErrNoSuchUser = errors.New("user with such id does not exist")
)