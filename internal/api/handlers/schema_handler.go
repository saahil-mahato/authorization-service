package handlers

import (
	"fmt"
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
	schemaService service.SchemaService
}

func NewHandler(schemaService service.SchemaService) *Handler {
	return &Handler{schemaService: schemaService}
}

func (h *Handler) WriteSchema(c *gin.Context) {
	var payload models.WriteSchemaPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedSchema, err := h.schemaService.WriteSchema(payload.TenantID, payload.Schema)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create authorization schema"})
		return
	}

	c.JSON(http.StatusCreated, savedSchema)
}

func (h *Handler) ListSchema(c *gin.Context) {
	tenantID := c.Param("id")
	continuousToken := c.Query("continuous_token")
	pageSizeStr := c.Query("page_size")

	// Convert pageSizeStr to uint32
	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	schemaList, err := h.schemaService.ListSchema(tenantID, continuousToken, uint32(pageSize))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch schema list for tenant %s", tenantID)})
		return
	}

	c.JSON(http.StatusOK, schemaList)
}

func (h *Handler) ReadSchema(c *gin.Context) {
	tenantID := c.Param("id")
	schemaVersion := c.Param("version")

	schema, err := h.schemaService.ReadSchema(tenantID, schemaVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch schema for tenant %s and schemaVersion %s", tenantID, schemaVersion)})
		return
	}

	c.JSON(http.StatusOK, schema)

}
