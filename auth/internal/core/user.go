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
}

func (s *ServiceImpl) GetUserInfo(ctx context.Context, initData string) (*UserInfo, Code, error) {
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
		if status.Code(err) == codes.InvalidArgument {
			return nil, CodeInvalidUserID, fmt.Errorf("invalid user id")
		}
		return nil, CodeInternal, err
	}

	if exist.Exists {
		return user, CodeOK, nil
	}

	_, err = s.usersClient.AddUser(ctx, &schema.AddUserRequest{
		Id: int64(user.Id),
		Firstname: user.Firstname,
		Username: user.Username,
	})
	if err != nil {
		return nil, CodeInternal, err
	}

	return user, CodeOK, nil
}
