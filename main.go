package main

import (
	"joke-web/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	config.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
