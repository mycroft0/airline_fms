package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/src/initializers"
	"server/src/models"
)

func ListAirline(c *gin.Context) {
	var airlineList []models.Charge

	initializers.DB.Table("airline").
		Limit(100).
		Scan(&airlineList)

	c.JSON(200, airlineList)
}

func AddCharges(c *gin.Context) {

	var requestData []models.Charge
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(requestData)

	// Save the data to the database
	if err := initializers.DB.Table("charge").Omit("name").Create(&requestData).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to save data to the database"})
		return
	}

	// Respond with a success message
	c.JSON(201, gin.H{"message": "Data saved successfully"})

}
