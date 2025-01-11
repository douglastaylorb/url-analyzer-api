// controllers/urlController.go

package controllers

import (
	"net/http"

	"github.com/douglastaylorb/url-analyzer-api/models"
	"github.com/douglastaylorb/url-analyzer-api/services"
	"github.com/gin-gonic/gin"
)

func AnalyzeURL(c *gin.Context) {
	var requestBody models.URLRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input inválido"})
		return
	}

	data, err := services.AnalyzeURL(requestBody.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao analisar a URL"})
		return
	}

	c.JSON(http.StatusOK, data)
}
