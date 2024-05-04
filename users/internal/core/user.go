package core

import (
	"context"
	"time"

	"github.com/Prokopevs/ccc/users/internal/model"
)

func (s *ServiceImpl) AddUser(ctx context.Context, user *model.UserReq) error {
	exists, err := s.db.IsUserWithIdExists(ctx, user.InviterId)
	if err != nil {
		return err
	}

	if !exists {
		user.InviterId = 0
	}
	now := time.Now()
	user.Createdat = &now

	return s.db.AddUser(ctx, user)
}

func (s *ServiceImpl) GetUser(ctx context.Context, id int) (*model.UserRes, bool, error) {
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

	return user, true, nil
}

func (s *ServiceImpl) IsUserWithIdExists(ctx context.Context, id int) (bool, error) {
	return s.db.IsUserWithIdExists(ctx, id)
}

func (s *ServiceImpl) GetUserReferrals(ctx context.Context, id int) ([]*model.UserReferrals, bool, error) {
	exists, err := s.db.IsUserWithIdExists(ctx, id)
	if err != nil {
		return nil, false, err
	}

	if !exists {
		return nil, true, ErrNoSuchUser
	}

	referrals, err := s.db.GetUserReferrals(ctx, id)
	if err != nil {
		return nil, false, err
	}

	return referrals, true, nil
}

