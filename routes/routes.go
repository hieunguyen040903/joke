package routes

import (
	"joke-web/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/joke", controllers.GetJoke(db))
	r.POST("/vote", controllers.PostVote(db))
}
