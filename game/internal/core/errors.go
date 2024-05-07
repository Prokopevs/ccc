package core

import "errors"

var (
	ErrNoSuchMultiplicator = errors.New("such multiplicator type doesn't exist")
	ErrNoSuchUser = errors.New("user with such id doesn't exist")
)