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

type SchemaHandler struct {
	schemaService service.SchemaService
}

func NewSchemaHandler(schemaService service.SchemaService) *SchemaHandler {
	return &SchemaHandler{schemaService: schemaService}
}

func (h *SchemaHandler) WriteSchema(c *gin.Context) {
	var payload models.WriteSchemaPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedSchema, err := h.schemaService.WriteSchema(payload.TenantID, payload.Schema)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create authorization schema. Error %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, savedSchema)
}

func (h *SchemaHandler) ListSchema(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch schema list for tenant %s. Error %s", tenantID, err.Error())})
		return
	}

	c.JSON(http.StatusOK, schemaList)
}

func (h *SchemaHandler) ReadSchema(c *gin.Context) {
	tenantID := c.Param("id")
	schemaVersion := c.Param("version")

	schema, err := h.schemaService.ReadSchema(tenantID, schemaVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch schema for tenant %s and schemaVersion %s. Error %s", tenantID, schemaVersion, err.Error())})
		return
	}

	c.JSON(http.StatusOK, schema)

}
