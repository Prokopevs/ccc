package server

import (
	"github.com/Prokopevs/ccc/schema"
	"github.com/Prokopevs/ccc/users/internal/core"
)

func convertPBAddUserToCore(user *schema.AddUserRequest) *core.UserReq {
	return &core.UserReq{
		Id:        int(user.Id),
		Firstname: user.Firstname,
		Username:  user.Username,
		InviterId: int(user.InviterId),
	}
}

func convertCoreUserToPB(user *core.UserRes) *schema.User {
	return &schema.User{
		Id:        int64(user.Id),
		Firstname: user.Firstname,
		Username:  user.Username,
		Referrals: user.Referrals,
	}
}
