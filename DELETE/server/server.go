package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteData struct {
	ID string `json:"id" binding:"required"`
}

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	// Initialize Gin Router
	router := gin.Default()

	// Define DELETE route
	router.DELETE("/delete", func(c *gin.Context) {
		var deleteData DeleteData

		// Bind JSON data to deleteData struct
		if err := c.ShouldBindJSON(&deleteData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Process the received data and send a response
		response := ResponseData{
			Message: "Successfully deleted record with ID: " + deleteData.ID,
		}
		c.JSON(http.StatusOK, response)
	})

	// Start the server on port 8080
	router.Run(":8080")
}
