package model

import "time"

type Match struct {
	Id           int       `json:"id"`
	MatchCatId   int       `json:"matchCatId"`
	UserCatId    int       `json:"userCatId"`
	UserId       int       `json:"userId"`
	TargetUserId int       `json:"targetUserId"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"createdAt"`
}
