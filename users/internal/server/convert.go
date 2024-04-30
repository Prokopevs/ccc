package server

import (
	"github.com/Prokopevs/ccc/schema"
	"github.com/Prokopevs/ccc/users/internal/core"
)

func convertPBAddUserToCore(user *schema.AddUserRequest) *core.User {
	return &core.User{
		Id:        int(user.Id),
		Firstname: user.Firstname,
		Username:  user.Username,
	}
}

func convertCoreUserToPB(user *core.User) *schema.User {
	return &schema.User{
		Id:        int64(user.Id),
		Firstname: user.Firstname,
		Username:  user.Username,
	}
}
