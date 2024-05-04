package model

import "time"

type Cats struct {
	Id          int       `json:"id"`
	User_id     string    `json:"user_id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int       `json:"ageInMonth"`
	Description string       `json:"description"`
	ImageUrls   []string  `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}


type PostCat struct {
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int       `json:"ageInMonth"`
	Description string    `json:"description"`
	ImageUrls   []string  `json:"imageUrls"`
}