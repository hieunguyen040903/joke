package main

import (
	"fmt"
	"joke-web/config"
	"joke-web/models"
	"joke-web/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	config.DB.AutoMigrate(&models.Vote{}, &models.Joke{})

	config.SeedData()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"API_URL": apiURL,
		})
	})

	routes.RegisterRoutes(r, db)

	r.Run(":8080")
}
