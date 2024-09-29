package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	// Initialize Gin Router
	router := gin.Default()

	// Define PUT route
	router.PUT("/update", func(c *gin.Context) {
		var updateData UpdateData

		// Bind JSON data to updateData struct
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Process the received data and send a response
		response := ResponseData{
			Message: "Updated: " + updateData.Name + "'s email to " + updateData.Email,
		}
		c.JSON(http.StatusOK, response)
	})

	// Start the server on port 8080
	router.Run(":8080")
}
