package handlers

import (
	"net/http"

	"github.com/AchrafIbrahim/WORK_APPLY_BOT/models"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/services"
	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	service *services.ApplicationService
}

func NewApplicationHandler(
	service *services.ApplicationService,
) *ApplicationHandler {
	return &ApplicationHandler{
		service: service,
	}
}

func (h *ApplicationHandler) CreateApplication(c *gin.Context) {
	var application models.Application

	if  err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ërror": "invalid request body",
		})
		return
	}
	
	if err := h.service.CreateApplication(
		c.Request.Context(),
		&application,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Application created successfully",
		"application": application,
	})
}