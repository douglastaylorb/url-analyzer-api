package routes

import (
	"github.com/douglastaylorb/url-analyzer-api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {

	router := gin.Default()

	//configuração de cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://net-kit.vercel.app",
			"https://netkit.douglastaylor.com.br",
		},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rotas URL Analyzer
	router.GET("/api/download-pdf", controllers.DownloadPDF)
	router.POST("/api/analyze", controllers.AnalyzeURL)

	// Rotas CPF Generator
	router.GET("/api/generate-cpf", controllers.GenerateCPFHandler)

	// Rotas MetaTag Validator
	router.POST("/api/validate-meta-tags", controllers.ValidateMetaTagHandler)

	// Rotas Port Finder
	router.POST("/api/scan-ports", controllers.ScanPortsController)

	// Rotas Data Interval Calculator
	router.POST("/api/data-difference", controllers.DataDifference)

	// Rotas QR Code Generator
	router.POST("/api/generate-qr", controllers.GenerateQR)

	// Rotas Password Generator
	router.POST("/api/generate-password", controllers.GeneratePasswordHandler)

	router.Run(":8080")
}
