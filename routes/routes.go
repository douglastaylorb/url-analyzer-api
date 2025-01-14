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

	router.Run(":8080")
}
