package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tommitoan/bazica"
	"time"
)

type BaziRequest struct {
	Date     string `json:"date" binding:"required"`     // Format: YYYY-MM-DD
	Time     string `json:"time" binding:"required"`     // Format: HH:mm
	Timezone string `json:"timezone" binding:"required"` // Example: Asia/Bangkok
	Gender   int    `json:"gender" binding:"required"`   // 0 = female, 1 = male
}

func calculateBazi(c *gin.Context) {
	// Parse the input JSON request
	var req BaziRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Parse location
	loc, err := time.LoadLocation(req.Timezone)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid timezone: %s", req.Timezone)})
		return
	}

	// Parse date and time
	datetime := fmt.Sprintf("%s %s", req.Date, req.Time)
	t, err := time.ParseInLocation("2006-01-02 15:04", datetime, loc)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date or time format"})
		return
	}

	// Calculate Ba-Zi chart
	chart, err := bazica.GetBaziChart(t, loc, req.Gender)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return the result as JSON
	c.JSON(200, chart)
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Define the API endpoint
	r.POST("/api/bazi", calculateBazi)

	// Start the server on port 8080
	r.Run(":8080")
}
