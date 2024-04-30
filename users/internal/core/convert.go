package core

import "github.com/Prokopevs/ccc/users/internal/pg"

func (n *User) toDB() *pg.User {
	return &pg.User{
		Id:        n.Id,
		Firstname: n.Firstname,
		Username:  n.Username,
	}
}

func convertDBUserToService(user *pg.User) *User {
	return &User{
		Id:        user.Id,
		Firstname: user.Firstname,
		Username:  user.Username,
	}
}