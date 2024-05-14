package server

import (
	"github.com/Prokopevs/ccc/game/internal/core"
	"github.com/Prokopevs/ccc/game/internal/model"
)

func validateScore(s model.Score) core.Code{
	if s.Id == 0 {
		return core.CodeIdCannotBeEmpty
	}

	if s.Score == 0 {
		return core.CodeScoreCannotBeEmpty
	}

	return core.CodeOK
}

func validateMultiplicator (s model.MultipUpdate) core.Code{
	if s.Id == 0 {
		return core.CodeIdCannotBeEmpty
	}

	if s.NameType == "" {
		return core.CodeMultiplicatorCannotBeEmpty
	}

	return core.CodeOK
}