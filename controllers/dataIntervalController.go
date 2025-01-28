package controllers

import (
	"net/http"

	"github.com/douglastaylorb/url-analyzer-api/services"
	"github.com/gin-gonic/gin"
)

func DataDifference(c *gin.Context) {
	var dates struct {
		StartDate string `json:"start_date" binding:"required"`
		EndDate   string `json:"end_date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&dates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	daysDifference, err := services.CalculateDateInterval(dates.StartDate, dates.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"days_difference": daysDifference})

}
