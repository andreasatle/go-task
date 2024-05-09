package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andreasatle/go-task/config"
	"github.com/andreasatle/go-task/database"
	"github.com/andreasatle/go-task/templates"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	database.ConnectDB(config.Database)
	r := gin.Default()

	r.Static("/static", "./static")

	// Create the templates, home made version of gin.LoadHTMLGlob
	tmpl := templates.CreateTemplates()

	// Define a route to render the index template
	r.GET("/public/contact/", func(c *gin.Context) {
		tmpl.Contact.Execute(c, gin.H{
			"title": "Gin Contact Web Page",
		})
	})

	// Define a route to render the index template
	r.GET("/public/resume/", func(c *gin.Context) {
		tmpl.Resume.Execute(c, gin.H{
			"title": "Gin Resume Web Page",
		})
	})

	// Define a route to render the index template
	r.GET("/public/", func(c *gin.Context) {
		tmpl.Index.Execute(c, gin.H{
			"title": "Gin Index Web Page",
		})

	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/public/home/")
	})

	// Start the server with TLS
	port := fmt.Sprintf(":%v", config.Server.Port)
	if err := http.ListenAndServeTLS(port, config.Tls.CertFile, config.Tls.KeyFile, r); err != nil {
		log.Fatalf("Certification error: %v", err)
	}
}
