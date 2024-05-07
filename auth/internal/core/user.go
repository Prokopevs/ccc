package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/Prokopevs/ccc/schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	CodeInvalidInitData Code = "INVALID_INITDATA"
	CodeInvalidUserID   Code = "INVALID_USER_ID"
)

type UserInfo struct {
	Id        int
	Firstname string
	Username  string
	InviterId int
}

type UserReferrals struct {
	ReferralId int
	Firstname  string
	Username   string
}

func (s *ServiceImpl) GetUserInfo(ctx context.Context, initData string, inviterId int) (*UserInfo, Code, error) {
	user, err := ValidateToken(initData, s.token)
	if err != nil {
		if errors.Is(err, errInitData) {
			return nil, CodeInvalidInitData, errInitData
		}
		return nil, CodeInternal, err
	}

	exist, err := s.usersClient.IsUserWithIdExists(ctx, &schema.IsUserWithIdExistsRequest{
		Id: int64(user.Id),
	})
	if err != nil {
		return nil, CodeInternal, err
	}

	if exist.Exists {
		return user, CodeOK, nil
	}

	_, err = s.usersClient.AddUser(ctx, &schema.AddUserRequest{
		Id:        int64(user.Id),
		Firstname: user.Firstname,
		Username:  user.Username,
		InviterId: int64(inviterId),
	})
	if err != nil {
		return nil, CodeInternal, err
	}

	return user, CodeOK, nil
}

func (s *ServiceImpl) GetUserReferrals(ctx context.Context, id int) ([]*UserReferrals, Code, error) {
	referrals, err := s.usersClient.GetUserReferrals(ctx, &schema.GetUserReferralsRequest{
		Id: int64(id),
	})
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return nil, CodeInvalidUserID, fmt.Errorf("invalid user id")
		}
		return nil, CodeInternal, err
	}

	return convertPBUserReferralsToUserReferrals(referrals.Referrals), CodeOK, nil
}
