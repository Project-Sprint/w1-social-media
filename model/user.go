package model

type User struct {
	Id       int    `json:"id", gorm:"text;not null"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
