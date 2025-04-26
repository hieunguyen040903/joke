package models

import "gorm.io/gorm"

type Vote struct {
	ID           int    `json:"id"`
	JokeID       int    `json:"joke_id"`
	UserCookieID string `json:"user_cookie_id"`
	VoteType     string `json:"vote_type"`
}

func SaveVote(db *gorm.DB, vote Vote) error {
	return db.Create(&vote).Error
}
