package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin Router
	router := gin.Default()

	// Define GET route
	router.GET("/data", func(c *gin.Context) {
		// Read query parameters from the GET request
		name := c.Query("name")
		email := c.Query("email")

		// Process the received data and send a response
		responseMessage := "Hello, " + name + "! We have received your email: " + email
		c.JSON(http.StatusOK, gin.H{"message": responseMessage})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
