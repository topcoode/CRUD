package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	// Initialize Gin Router
	router := gin.Default()

	// Define POST route
	router.POST("/submit", func(c *gin.Context) {
		var requestData RequestData

		// Bind JSON data to requestData struct
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Process the received data and send a response
		response := ResponseData{
			Message: "Hello, " + requestData.Name + "! We have received your email: " + requestData.Email,
		}
		c.JSON(http.StatusOK, response)
	})

	// Start the server on port 8080
	router.Run(":8080")
}
