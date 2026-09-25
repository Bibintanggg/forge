package handler

import (
	"net/http"

	"github.com/Bibintanggg/forge/internal/api/service"
	"github.com/Bibintanggg/forge/internal/doctor"
	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	service service.DoctorService
}

func NewDoctorHandler(service service.DoctorService) *DoctorHandler {
	return &DoctorHandler{
		service: service,
	}
}

func (h *DoctorHandler) GetDoctor(c *gin.Context) {
	result := doctor.Run()

	c.JSON(http.StatusOK, result)
}
