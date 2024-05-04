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
			ReferralId: ref.ReferralId,
			Firstname:  ref.Firstname,
			Username:   ref.Username,
		})
	}
	return updatedUserReferrals
}
