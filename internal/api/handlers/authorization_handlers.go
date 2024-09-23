package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saahil-mahato/authorization-service/internal/models"
	"github.com/saahil-mahato/authorization-service/internal/service"
)

const (
	INVALID_ID = "Invalid ID"
)

type Handler struct {
	authService service.AuthorizationService
}

func NewHandler(authService service.AuthorizationService) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) CreateAuth(c *gin.Context) {
	var config models.AuthorizationConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.CreateConfig(&config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create authorization config"})
		return
	}

	c.JSON(http.StatusCreated, config)
}

func (h *Handler) GetAuth(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": INVALID_ID})
		return
	}

	config, err := h.authService.GetConfig(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Authorization config not found"})
		return
	}

	c.JSON(http.StatusOK, config)
}

func (h *Handler) ListAuth(c *gin.Context) {
	configs, err := h.authService.ListConfigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list authorization configs"})
		return
	}

	c.JSON(http.StatusOK, configs)
}

func (h *Handler) UpdateAuth(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": INVALID_ID})
		return
	}

	var config models.AuthorizationConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.ID = uint(id)

	if err := h.authService.UpdateConfig(&config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update authorization config"})
		return
	}

	c.JSON(http.StatusOK, config)
}

func (h *Handler) DeleteAuth(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": INVALID_ID})
		return
	}

	if err := h.authService.DeleteConfig(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete authorization config"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authorization config deleted successfully"})
}

func (h *Handler) CheckAuthorization(c *gin.Context) {
	var request struct {
		URL      string `json:"url"`
		Resource string `json:"resource"`
		Role     string `json:"role"`
		Action   string `json:"action"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorized, prismaSchema, err := h.authService.CheckAuthorization(request.URL, request.Resource, request.Role, request.Action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check authorization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authorized": authorized, "prismaSchema": prismaSchema})
}
