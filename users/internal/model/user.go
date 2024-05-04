package model

import "time"

type UserReq struct {
	Id        int
	Firstname string
	Username  string
	InviterId int
	Createdat *time.Time
}

type UserRes struct { 
	Id        int        `db:"id,omitempty"`
	Firstname string     `db:"firstname,omitempty"`
	Username  string     `db:"username,omitempty"`
	Createdat *time.Time `db:"createdat,omitempty"`
}

type UserReferrals struct {
    ReferralId int64    `db:"referralid,omitempty"`
    Firstname  string `db:"firstname,omitempty"`
    Username   string `db:"username,omitempty"`
}