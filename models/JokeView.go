package models

import "gorm.io/gorm"

type JokeView struct {
	gorm.Model
	UserCookieID string `gorm:"index"`
	JokeID       uint   `gorm:"index"`
}
