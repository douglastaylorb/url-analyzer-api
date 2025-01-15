package controllers

import (
	"net/http"

	"github.com/douglastaylorb/url-analyzer-api/models"
	"github.com/douglastaylorb/url-analyzer-api/services"
	"github.com/gin-gonic/gin"
)

func ValidateMetaTagHandler(c *gin.Context) {
	var json models.MetaTagRequest
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input inválido"})
		return
	}

	metaTags, err := services.ValidateMetaTags(json.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao validar meta tags"})
		return
	}

	c.JSON(http.StatusOK, metaTags)

}
