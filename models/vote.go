package models

import (
	"errors"

	"gorm.io/gorm"
)

type Vote struct {
	ID           int    `json:"id"`
	JokeID       int    `json:"joke_id"`
	UserCookieID string `json:"user_cookie_id"`
	IsLike       bool   `gorm:"not null"`
}

func SaveVote(db *gorm.DB, vote Vote) error {
	var existingVote Vote
	if err := db.Where("user_cookie_id = ? AND joke_id = ?", vote.UserCookieID, vote.JokeID).First(&existingVote).Error; err == nil {
		return errors.New("you have already voted for this joke")
	}

	return db.Create(&vote).Error
}
