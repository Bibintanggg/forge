package api

import (
	"net/http"

	"github.com/Bibintanggg/forge/internal/api/handler"
	"github.com/Bibintanggg/forge/internal/api/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	doctorService := service.NewDoctorService()
	doctorHandler := handler.NewDoctorHandler(*doctorService)

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "forge-api",
		})
	})

	router.GET("/api/doctor", doctorHandler.GetDoctor)

	return router
}
