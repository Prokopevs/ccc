package core

import "github.com/Prokopevs/ccc/schema"

type ServiceImpl struct {
	usersClient schema.UsersClient
	token       string
}

func NewServiceImpl(usersClient schema.UsersClient, token string) *ServiceImpl {
	return &ServiceImpl{
		usersClient: usersClient,
		token:       token,
	}
}
