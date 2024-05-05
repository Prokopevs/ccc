package core

import "github.com/Prokopevs/ccc/schema"

func convertUserTelegramToUserInfo(u *UserTelegram) *UserInfo {
	return &UserInfo{
		Id:        u.Id,
		Firstname: u.Firstname,
		Username:  u.Username,
	}
}

func convertPBUserReferralsToUserReferrals(referrals []*schema.Referrals) []*UserReferrals {
	var updatedUserReferrals []*UserReferrals
	for _, ref := range referrals {
		updatedUserReferrals = append(updatedUserReferrals, &UserReferrals{
			ReferralId: int(ref.ReferralId),
			Firstname:  ref.Firstname,
			Username:   ref.Username,
		})
	}
	return updatedUserReferrals
}