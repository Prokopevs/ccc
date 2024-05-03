package core

import (
	"context"
	"time"
)

type UserReq struct {
	Id        int     
	Firstname string     
	Username  string     
	ReferralId int
}

type UserRes struct {
	Id        int     
	Firstname string     
	Username  string     
	Referrals []int64
}

func (s *ServiceImpl) AddUser(ctx context.Context, user *UserReq) error {
	exists, err := s.db.IsUserWithIdExists(ctx, user.ReferralId)
	if err != nil {
		return err
	}

	u := user.toDB()
	if exists {
		u.ReferralId = user.Id
	} else {
		u.ReferralId = 0
	}
	now := time.Now()
	u.Createdat = &now

	return s.db.AddUser(ctx, u)
}

func (s *ServiceImpl) GetUser(ctx context.Context, id int) (*UserRes, bool, error) {
	exists, err := s.db.IsUserWithIdExists(ctx, id)
	if err != nil {
		return nil, false, err
	}

	if !exists {
		return nil, true, ErrNoSuchUser
	}

	user, err := s.db.GetUser(ctx, id)
	if err != nil {
		return nil, false, err
	}

	return convertDBUserToService(user), true, nil
}

func (s *ServiceImpl) IsUserWithIdExists(ctx context.Context, id int) (bool, error) {
	return s.db.IsUserWithIdExists(ctx, id)
}