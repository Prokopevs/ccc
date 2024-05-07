package pg

import "errors"

var (
	ErrNoEnoughScore = errors.New("not enough score")
	ErrMaxLevel = errors.New("not more level")
)