package core

import (
	"encoding/json"
	"errors"
	"net/url"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

var (
	errInitData = errors.New("invalid initData provided")
)

func ValidateToken(initData, token string) (UserInfo, error) {
	initData, err := url.QueryUnescape(initData)
	if err != nil {
		return UserInfo{}, err
	}

	// check telegram string
	err = initdata.Validate(initData, token, 0)
	if err != nil {
		return UserInfo{}, errInitData
	}

	// get user data
	q, _ := url.ParseQuery(initData)
	var user UserInfo
	err = json.Unmarshal([]byte(q["user"][0]), &user)
	if err != nil {
		return UserInfo{}, err
	}

	return user, nil
}