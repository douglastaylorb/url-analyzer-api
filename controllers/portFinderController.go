package controllers

import (
	"net/http"
	"strings"

	"github.com/douglastaylorb/url-analyzer-api/services"
	"github.com/gin-gonic/gin"
)

// ScanPortsController processa a requisição HTTP para varreduras de portas.
// Ele espera um corpo de requisição em JSON que especifica o domínio.
// O controlador valida, maneja erros, executa a varredura e devolve resultados.
func ScanPortsController(c *gin.Context) {
	var json struct {
		Domain     string `json:"domain"`
		PortString string `json:"ports"` // novo campo esperado JSON
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request data is invalid"})
		return
	}

	domain := strings.TrimSpace(json.Domain)
	ports, err := services.ParsePorts(json.PortString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	openPorts := services.ScanOpenPorts(domain, ports)

	c.JSON(http.StatusOK, gin.H{
		"open_ports": openPorts,
	})
}
