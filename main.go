package main

import (
	"fmt"
	"joke-web/config"
	"joke-web/models"
	"joke-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	db, err := config.ConnectDB()

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	config.DB.AutoMigrate(&models.Vote{}, &models.Joke{})

	config.SeedData()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	routes.RegisterRoutes(r, db)

	r.Run(":8080")
}
