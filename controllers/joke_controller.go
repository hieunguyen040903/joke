package controllers

import (
	"joke-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetJoke(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var joke models.Joke
		if err := db.Where("id NOT IN (?)", db.Model(&models.Vote{}).Select("joke_id")).First(&joke).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "That's all the jokes for today! Come back another day!"})
			return
		}
		c.JSON(http.StatusOK, joke)
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
