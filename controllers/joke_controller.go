package controllers

import (
	"errors"
	"joke-web/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrNoJokesLeft = errors.New("that's all the jokes for today! Come back another day!")

func GetJoke(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCookieID := c.DefaultQuery("user_cookie_id", "")
		if userCookieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No user_cookie_id provided"})
			return
		}

		var joke models.Joke
		err := db.Table("jokes").Where("id NOT IN (SELECT joke_id FROM joke_views WHERE user_cookie_id = ?)", userCookieID).
			First(&joke).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusOK, gin.H{"message": ErrNoJokesLeft.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		jokeView := models.JokeView{
			UserCookieID: userCookieID,
			JokeID:       joke.ID,
		}

		err = db.Create(&jokeView).Error
		if err != nil {
			log.Printf("Error saving joke view: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save joke view"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"joke": joke,
		})
	}
}

func PostVote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vote models.Vote

		if err := c.ShouldBindJSON(&vote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vote"})
			return
		}

		var existingVote models.Vote
		if err := db.Where("user_cookie_id = ? AND joke_id = ?", vote.UserCookieID, vote.JokeID).First(&existingVote).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "you have already voted for this joke"})
			return
		}

		if err := db.Create(&vote).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save vote"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "vote saved successfully"})
	}
}
