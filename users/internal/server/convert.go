package server

import (
	"github.com/Prokopevs/ccc/schema"
	"github.com/Prokopevs/ccc/users/internal/model"
)

func convertPBAddUserToCore(user *schema.AddUserRequest) *model.UserReq {
	return &model.UserReq{
		Id:        int(user.Id),
		Firstname: user.Firstname,
		Username:  user.Username,
		InviterId: int(user.InviterId),
	}
}

func convertCoreUserToPB(user *model.UserRes) *schema.User {
	return &schema.User{
		Id:        int64(user.Id),
		Firstname: user.Firstname,
		Username:  user.Username,
	}
}

func convertCoreReferralsToPB(referrals []*model.UserReferrals) []*schema.Referrals {
	var updatedUserReferrals []*schema.Referrals
	for _, ref := range referrals {
		updatedUserReferrals = append(updatedUserReferrals, &schema.Referrals{
			ReferralId: int64(ref.ReferralId),
			Firstname:  ref.Firstname,
			Username:   ref.Username,
		})
	}
	return updatedUserReferrals
}

func convertCoreUsersToPB(users []*model.UserRes) []*schema.User {
	var updatedUsers []*schema.User
	for _, usr := range users {
		updatedUsers = append(updatedUsers, &schema.User{
			Id:        int64(usr.Id),
			Firstname: usr.Firstname,
			Username:  usr.Username,
		})
	}
	return updatedUsers
}
