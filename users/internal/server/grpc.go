package server

import (
	"github.com/Prokopevs/ccc/schema"
	"github.com/Prokopevs/ccc/users/internal/core"
)

type GRPC struct {
	schema.UnimplementedUsersServer

	service core.Service
}

func NewGRPC(usersService core.Service) *GRPC {
	return &GRPC{
		service: usersService,
	}
}
