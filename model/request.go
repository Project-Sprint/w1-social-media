package model

type RequestMatch struct {
	MatchCatId int    `json:"matchCatid"`
	UserCatId  int    `json:"userCatId"`
	Message    string `json:"message"`
}
