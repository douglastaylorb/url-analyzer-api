package controllers

import (
	"net/http"

	"github.com/douglastaylorb/url-analyzer-api/models"
	"github.com/douglastaylorb/url-analyzer-api/services"
	"github.com/gin-gonic/gin"
)

func GenerateCPFHandler(c *gin.Context) {
	cpf := services.GenerateValidCPF()
	response := models.CPFResponse{CPF: cpf}

	c.JSON(http.StatusOK, response)

}
