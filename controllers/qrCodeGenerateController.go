package controllers

import (
	"io"
	"net/http"

	"github.com/douglastaylorb/url-analyzer-api/services"
	"github.com/gin-gonic/gin"
)

func GenerateQR(c *gin.Context) {
	url := c.PostForm("url")
	includeLogo := c.PostForm("includeLogo") == "true"

	var logoData []byte
	if includeLogo {
		file, _ := c.FormFile("logo")
		if file != nil {
			f, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível abrir o arquivo de logo"})
				return
			}
			defer f.Close()

			logoData, err = io.ReadAll(f)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler o arquivo de logo"})
				return
			}
		}
	}

	qrCode, err := services.GenerateQR(url, includeLogo, logoData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar QR code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"qr": qrCode})
}
