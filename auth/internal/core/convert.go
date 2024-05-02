package core 

func convertUserTelegramToUserInfo(u *UserTelegram) *UserInfo {
	return &UserInfo{
		Id:        u.Id,
		Firstname: u.Firstname,
		Username:  u.Username,
	}
}