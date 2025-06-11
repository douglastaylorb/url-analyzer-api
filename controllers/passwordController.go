package controllers

import (
	"net/http"

	"github.com/douglastaylorb/url-analyzer-api/models"
	"github.com/douglastaylorb/url-analyzer-api/services"
	"github.com/gin-gonic/gin"
)

func GeneratePasswordHandler(c *gin.Context) {
	var request models.PasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos"})
		return
	}

	response, err := services.GeneratePassword(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
