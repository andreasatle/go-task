package main

import (
	"net/http"

	"github.com/andreasatle/go-task/config"
	"github.com/andreasatle/go-task/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	database.ConnectDB(config.Database)
	r := gin.Default()

	r.Static("/static", "./static")
	// Load templates from the "templates" directory
	r.LoadHTMLGlob("templates/*.tmpl")

	// Define a route to render the index template
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Gin Index Web Page",
		})
	})

	// Define a route to render the index template
	r.GET("/resume", func(c *gin.Context) {
		c.HTML(http.StatusOK, "resume.tmpl", gin.H{
			"title": "Gin Resume Web Page",
		})
	})

	// Start the server
	r.Run(":8080")
}
