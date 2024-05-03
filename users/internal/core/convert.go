package core

import (
	"github.com/Prokopevs/ccc/users/internal/pg"
)

func (n *UserReq) toDB() *pg.UserReq {
	return &pg.UserReq{
		Id:        n.Id,
		Firstname: n.Firstname,
		Username:  n.Username,
	}
}

func convertDBUserToService(user *pg.UserRes) *UserRes {
	return &UserRes{
		Id:        user.Id,
		Firstname: user.Firstname,
		Username:  user.Username,
		Referrals: user.Referrals,
	}
}