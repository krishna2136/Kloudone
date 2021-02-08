package main

import (
	"github.com/IBM-Bluemix/cf_deployment_tracker_client_go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {  
	cf_deployment_tracker.Track()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome all",
		})
	})

	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "happy new year",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
