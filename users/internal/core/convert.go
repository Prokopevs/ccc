package core

import (
	"encoding/binary"

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
	intSlice := make([]int64, len(user.Referrals))

	for i := 0; i < len(user.Referrals); i += 8 {
		bits := binary.LittleEndian.Uint64(user.Referrals[i : i+8])
		intSlice = append(intSlice, int64(bits))
	}
	
	return &UserRes{
		Id:        user.Id,
		Firstname: user.Firstname,
		Username:  user.Username,
		Referrals: intSlice,
	}
}