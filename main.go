package main

import (
	"joke-web/config"
	"joke-web/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Vote{}, &models.Joke{})
	config.SeedData()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
