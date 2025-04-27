package main

import (
	"fmt"
	"joke-web/config"
	"joke-web/models"
	"joke-web/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080"
	}

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	db.AutoMigrate(&models.Vote{}, &models.Joke{})
	config.SeedData()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"API_URL": apiURL,
		})
	})

	routes.RegisterRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
